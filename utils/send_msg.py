#!/usr/bin/env python
# -*- coding: utf-8 -*-

import os
import sys

import boto3
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

print(queue)

msg_num = 4
msg_list = [
    {
        'Id' : '{}'.format(i+1),
        'MessageBody' : 'msg_{}'.format(i+1)
    } for i in range(msg_num)]
response = queue.send_messages(Entries=msg_list)
print(response)
