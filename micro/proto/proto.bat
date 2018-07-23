@echo off
echo ----------------------------------------------------------
echo                      Welcome Use     
echo ----------------------------------------------------------
echo working .....
@echo on
protoc --go_out=plugins=micro:. micro.proto
@echo off
echo.