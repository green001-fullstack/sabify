import chromadb

chroma_client = chromadb.Client()

collection_name = "test_collection"

collection = chroma_client.get_or_create_collection(collection_name)

documents = [
    {"id": "doc1", "text": "Hello, world!"},
    {"id": "doc2", "text": "Amazon EC2 is a virtual server in the AWS Cloud."},
    {"id": "doc3", "text": "Amazon S3 is an object storage service used to store files and data."}
]

for doc in documents:
    collection.upsert(ids=doc["id"], documents=[doc["text"]])

query = "Hello, world!"

results = collection.query(
    query_texts=[query],
    n_results=3,
)

print(results)