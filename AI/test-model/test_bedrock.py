from dotenv import load_dotenv
from openai import OpenAI

load_dotenv()

client = OpenAI()

response = client.responses.create(
    model="openai.gpt-oss-20b",
    input="Explain Amazon EC2 in simple terms."
)

print(response.output_text)