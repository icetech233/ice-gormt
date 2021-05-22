@echo off

del /f /s /q model

gormt -H=10.200.37.201 -d=batrix_rds -p=cZtv@mz-21030917 -u=rhpb --port=33060

rem gormt -H=127.0.0.1 -d=batrix_rds -p=cZtv@mz-21030917 -u=rhpb --port=3306

rem gormt -H=127.0.0.1 -d=zone_data_1 -p=123456 -u=root --port=3306

