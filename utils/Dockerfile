#PythonのDockerイメージを指定
FROM python:3.6

#ファイルをappディレクトリに追加
COPY . /app/

#ルートディレクトリ設定
WORKDIR /app

RUN pip install boto3
RUN pip install requests
#コマンド実行
ENTRYPOINT ["python", "script.py"]