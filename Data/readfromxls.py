import xlrd
import pymysql

db = pymysql.connect("localhost", "root", "password", "CanEat")
cursor = db.cursor()

book = xlrd.open_workbook("2017.xls")
sheet = book.sheet_by_index(0)

for i in range(2, 1225):
    row = sheet.row_values(i)
    try:
        cursor.execute("INSERT INTO foods VALUES(%d, '%s', %d, %d, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f)" % (
            i-1,
            row[1],
            100 if row[2] == '' else int(row[2]),
            int(row[3]),
            round(row[4],2),
            round(row[5],2),
            round(row[6],2),
            round(row[7],2),
            round(row[8],2),
            round(row[9],2),
            round(row[10],2),
            round(row[11],2),
            round(row[12],2),
            round(row[18],2),
            round(row[13],2),
            round(row[14],2),
            round(row[15],2),
            round(row[16],2),
            round(row[19],2),
            ))
        db.commit()
    except:
        db.rollback()

db.close()
