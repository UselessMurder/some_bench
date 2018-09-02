import time

file = open("./bin", "rb")

content = file.read()

file.close()

start = time.time()

hstr = content.hex()

#print(hstr)

content = bytes.fromhex(hstr)

end = time.time()

print(end - start, "s.")


