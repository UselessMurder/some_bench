import time
import msg_pb2

file = open("./bin", "rb")

content = file.read()

file.close()

message = msg_pb2.msg()
message.name = "Hello"
message.id = 0
message.content = content

start = time.time()

data = message.SerializeToString()
new_message = msg_pb2.msg()
new_message.ParseFromString(data)

end = time.time()

print(end - start, "s.")
