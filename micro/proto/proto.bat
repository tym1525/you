@echo off
echo ----------------------------------------------------------
echo                      Welcome Use     
echo ----------------------------------------------------------
echo working .....
@echo on
protoc --proto_path=. --micro_out=. --go_out=. micro.proto
protoc --go_out=plugins=micro:. micro.proto
@echo off
echo.