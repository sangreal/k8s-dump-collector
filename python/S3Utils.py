import boto3
from botocore import config
import logging


def write_to_s3(bucket: str, keyname: str, file_path: str):
    try:
        s3 = boto3.client('s3', 'ap-southeast-1', config=config.Config(s3={'addressing_style': 'path'}))
        print("begin to write to S3 ： " + bucket
                     + "  key : " + keyname
                     + " file_path : " + file_path)
        s3.upload_file(file_path, bucket, keyname)
    except Exception as e:
        print("write to S3 error : " + str(e))


