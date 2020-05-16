# #!/usr/bin/env python
# # -*- coding: utf-8 -*-
import os
import sys

import boto3
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
            print(message.body)
            message.delete()
    else:
        # メッセージがなくなったらbreak
        print("No remaining message")
        break
