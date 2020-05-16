# #!/usr/bin/env python
# # -*- coding: utf-8 -*-
import os
import sys
import subprocess
import json

import boto3
import requests



def execCode(code, input, expectOutput):
    path = './code.py'
    
    with open(path, mode='w') as f:
        f.write(code)
    
    proc = subprocess.run("python code.py", shell=True,input=input, encoding='utf-8', stdout=subprocess.PIPE)
    output = proc.stdout.rstrip()

    print(output)
    print(expectOutput)
    if(output == expectOutput):
        return "AC", output
    else:
        return "WA", output
    


def main():
    client = boto3.resource('sqs',
                        endpoint_url=os.getenv("SQS_ENDPOINT"),
                        region_name=os.getenv("AWS_REGION"),
                        aws_secret_access_key=os.getenv("AWS_ACCESS_KEY"),
                        aws_access_key_id=os.getenv("AWS_SECRET_KEY"),
                        use_ssl=False)
    queue = client.get_queue_by_name(QueueName='python')
    print(queue)
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
                type Answer struct {
                    ID         int       `json:"id"`
                    QuestionID int       `json:"question_id"`
                    Answer     string    `json:"answer"`
                    Status     string    `json:"status"`
                    Result     string    `json:"result"`
                    Detail     string    `json:"detail"`
                    CreatedAt  time.Time `json:"created_at"`
                    UpdatedAt  time.Time `json:"updated_at"`
                }
                """
                code = json.loads(message.body)

                result, output = execCode(code['answer'], '1 2 3', 'No')
                code['result'] = result
                code['detail'] = output
                url = os.getenv("SERVER_URL") + "answers/" + str(code['id'])
                print(url)
                print(code)
                res = requests.put(url, json.dumps(code))
                print(res)
                message.delete()


        else:
            # メッセージがなくなったらbreak
            print("No remaining message")
            break


if __name__ == "__main__":
    main()



# curl localhost:8080/answers/1 -X POST -d '{"question_id":1,"answer":"A, B, C = input().split()\\nA = int(A)\\nB = int(B)\\nC = int(C)\\nif A < C < B:\\n\\tprint(\'Yes\')\\nelif B < C < A:\\n\\tprint(\'Yes\')\\nelse:\\n\\tprint(\'No\')","status":"status","result":"result","detail":"detail"}'