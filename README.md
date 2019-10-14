# testgoapi
An example of the rest api (CURD) by golang and Microsoft SQL SERVER

This an example of golang script works as the back-end in the server.
The environments setting and tools are below

1 - Install MS-SQL SERVER EXPRESS 2005 or newly released versions
2 - Generate the database for MS-SQL SERVER by SQL script as below 

USE [bookshop]
GO
/****** Object:  Table [dbo].[books]    Script Date: 10/14/2019 9:52:26 AM ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[books](
	[isbn] [int] IDENTITY(1,1) NOT NULL,
	[bookname] [varchar](max) NULL,
	[bookprice] [int] NULL
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO

3- Enable SQL Server Authentication in MS-SQL Server with user "sa" and password ="your password"    
3- Install go 1.xx 
4- Install Postman
