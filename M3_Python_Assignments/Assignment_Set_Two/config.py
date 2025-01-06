import os

class Config:
    SECRET_KEY = os.environ.get('SECRET KEY') or "you-will-never-guess"
    SQLALCHEMY_DATABASE_URI= r'sqlite:///C:\AdityaPersonal\nexTurn\AdityaSharma_NexTurn_Assignments\M3_Python_Assignments\Assignment_3\db\mydb.db'
    SQLALCHEMY_TRACK_MODIFICATIONS = False