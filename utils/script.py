# #!/usr/bin/env python
# # -*- coding: utf-8 -*-
import os
import sys
import subprocess
import json
from random import randint
from time import sleep

import boto3
import requests


"""
type TestCase struct {
	gorm.Model
	QuestionID int
	Input      string
	Output     string
}

{'ID': 1, 'CreatedAt': '2020-05-16T12:27:33.630985Z', 'UpdatedAt': '2020-05-16T12:27:33.630985Z', 'DeletedAt': None, 'QuestionID': 1, 'Input': '10 20', 'Output': '30'}

"""
def execCode(code, testcases):
    path = './code.py'
    
    with open(path, mode='w') as f:
        f.write(code)

    cnt = 0
    ac = 0
    err = []
    for i in testcases:
        print(i)
        cnt = cnt + 1
        proc = subprocess.run("python code.py", shell=True,input=i['Input'], encoding='utf-8', stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        output = proc.stdout.rstrip()

        print(output)
        print(i['Output'])
        if(output == i['Output']):
            ac = ac + 1
        else:
            err.append(proc.stderr.rstrip())
    
    result = ""
    for l in list(set(err)):
        result = result + l
     
    return cnt, ac, result
    


def main():
    client = boto3.resource('sqs',
        endpoint_url=os.getenv("SQS_ENDPOINT"),
        region_name=os.getenv("AWS_REGION"),
        aws_secret_access_key=os.getenv("AWS_ACCESS_KEY"),
        aws_access_key_id=os.getenv("AWS_SECRET_KEY"),
        use_ssl=False)

    try:
        # キューの名前を指定してインスタンスを取得
        queue = client.get_queue_by_name(QueueName='python')
    except Exception as e:
        # 指定したキューがない場合はexception
        print(e)
        sys.exit(-1)

    while True:
        # メッセージを取得
        msg_list = queue.receive_messages(MaxNumberOfMessages=10)
        if msg_list:
            for message in msg_list:
                """
                type SQSData struct {
                    gorm.Model
                    AnswerID  uint   `json:"answer_id"`
                    Answer    string `json:"answer"`
                    TestCases []TestCase
                }
                """
                code = json.loads(message.body)
                print(code)
                cnt, ac, err = execCode(code['answer'], code['TestCases'])
                if(err == ""):
                    code['result'] = "AC" if cnt == ac else "WA"
                else:
                    code['result'] = "RE"
                
                detail = str(ac) + "/" + str(cnt)
                code['detail'] = detail.rstrip()
                code['error'] = err

                url = os.getenv("SERVER_URL") + "answers/" + str(code['answer_id'])
                res = requests.put(url, json.dumps(code))
                print(res)
                
                message.delete()


        else:
            # メッセージがなくなったらbreak
            print("No remaining message")
            time = randint(1,10)
            print("wait " + str(time) + " sec")
            sleep(time)
            


if __name__ == "__main__":
    main()



# curl localhost:8080/answers/1 -X POST -d '{"question_id":1,"answer":"A, B, C = input().split()\\nA = int(A)\\nB = int(B)\\nC = int(C)\\nif A < C < B:\\n\\tprint(\'Yes\')\\nelif B < C < A:\\n\\tprint(\'Yes\')\\nelse:\\n\\tprint(\'No\')","status":"status","result":"result","detail":"detail"}'