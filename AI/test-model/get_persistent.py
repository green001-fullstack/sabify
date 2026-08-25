import chromadb

client = chromadb.PersistentClient(path= "./chroma_db")

name = "vehicles"

collection = client.get_collection(name)

print(collection.name)
print(collection.get())