Steps:

1) Get the path of the file
2) Read the file from provided path ---> It returns byte format of YAML file
3) Convert the YAML byte into JSON byte
4) Now, fill the JSON value in proper struct.

NOTE: Whenever dealing with reading and writting of file make sure everything is in bytes format. It should not be in string or int, et.

