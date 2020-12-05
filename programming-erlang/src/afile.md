c(afile_server).

c(afile_client).

FileServer = afile_server:start(".").

afile_client:ls(FileServer).

afile_client:get_file(FileServer, "hello.erl").

afile_client:put_file(S, "abc.txt", "hello abc").    

afile_client:get_file(S, "afile.md").