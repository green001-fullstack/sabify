import json
import os
import tempfile

import boto3
from dotenv import load_dotenv

from langchain_chroma import Chroma
from langchain_core.documents import Document
from langchain_core.embeddings import Embeddings


# Load variables from .env
load_dotenv()


class TitanV2Embeddings(Embeddings):

    def __init__(self):
        self.client = boto3.client(
            "bedrock-runtime",
            region_name="us-east-1"
        )

        self.model_id = "amazon.titan-embed-text-v2:0"

    def embed_documents(self, texts: list[str]) -> list[list[float]]:

        embeddings = []

        for text in texts:

            response = self.client.invoke_model(
                modelId=self.model_id,
                body=json.dumps({
                    "inputText": text,
                    "dimensions": 1024,
                    "normalize": True
                })
            )

            response_body = json.loads(
                response["body"].read()
            )

            embeddings.append(
                response_body["embedding"]
            )

        return embeddings

    def embed_query(self, text: str) -> list[float]:

        response = self.client.invoke_model(
            modelId=self.model_id,
            body=json.dumps({
                "inputText": text,
                "dimensions": 1024,
                "normalize": True
            })
        )

        response_body = json.loads(
            response["body"].read()
        )

        return response_body["embedding"]


# Create the embedding model
embeddings_model = TitanV2Embeddings()


# Sample documents
SAMPLE_DOCS = [
    Document(
        page_content=(
            "A vector store stores numerical vector representations "
            "of data."
        ),
        metadata={
            "source": "vector_stores.txt"
        }
    ),

    Document(
        page_content=(
            "Embeddings convert text into numerical vectors "
            "that capture semantic meaning."
        ),
        metadata={
            "source": "embeddings.txt"
        }
    ),

    Document(
        page_content=(
            "ChromaDB is a vector database that stores embeddings "
            "and performs similarity searches."
        ),
        metadata={
            "source": "chromadb.txt"
        }
    ),

    Document(
        page_content=(
            "Amazon EC2 provides resizable virtual servers "
            "in the AWS Cloud."
        ),
        metadata={
            "source": "aws.txt"
        }
    ),
]


def similarity_search_with_scores():

    with tempfile.TemporaryDirectory() as tmpdir:

        vectorstore = Chroma.from_documents(
            documents=SAMPLE_DOCS,
            embedding=embeddings_model,
            persist_directory=tmpdir
        )

        query = "Explain vector stores."

        results_with_scores = (
            vectorstore.similarity_search_with_score(
                query,
                k=3
            )
        )

        print(
            f"Top 3 results for: '{query}'\n"
        )

        for i, (doc, score) in enumerate(
            results_with_scores
        ):

            print(f"Result {i + 1}")
            print(f"Content: {doc.page_content}")
            print(f"Score: {score:.4f}")
            print(
                f"Source: "
                f"{doc.metadata.get('source', 'Unknown')}"
            )
            print("-" * 50)


similarity_search_with_scores()