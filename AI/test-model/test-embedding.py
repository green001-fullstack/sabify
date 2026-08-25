import os
import json
import boto3

from dotenv import load_dotenv


load_dotenv()

# Take your existing Bedrock API key
# from OPENAI_API_KEY and make it available
# to Boto3 using AWS's expected variable name.
os.environ["AWS_BEARER_TOKEN_BEDROCK"] = os.getenv(
    "OPENAI_API_KEY"
)


client = boto3.client(
    "bedrock-runtime",
    region_name="us-east-1"
)


response = client.invoke_model(
    modelId="amazon.titan-embed-text-v2:0",
    body=json.dumps({
        "inputText": "Amazon EC2 is a virtual server in AWS.",
        "dimensions": 1024,
        "normalize": True
    })
)


response_body = json.loads(
    response["body"].read()
)


embedding = response_body["embedding"]


print("Embedding generated successfully!")
print("Dimension:", len(embedding))
print("First 10 values:", embedding[:10])