# README

## Create this module

~~~
go mod init github.com/xralf/dynamo/go/guregu
go mod tidy
~~~

## Run the program

~~~
go run example.go
~~~

## Expected example output

The code uses a ramdom number generator for `count`, so the output will vary.

~~~
guregu(master ✗) go run example.go 
Table dropped successfully!
Table created successfully!
Last item inserted
        {613 2021-09-10 14:28:28.72443 -0700 PDT hello 1 [{101 2021-09-10 14:28:28.724429 -0700 PDT aaa 4 [] [] map[]  } {102 2021-09-10 14:28:28.72443 -0700 PDT bbb 4 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
All items:
0:      {613 2021-09-10 14:28:28.587935 -0700 PDT hello 7 [{101 2021-09-10 14:28:28.587934 -0700 PDT aaa 8 [] [] map[]  } {102 2021-09-10 14:28:28.587934 -0700 PDT bbb 3 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
1:      {613 2021-09-10 14:28:28.604661 -0700 PDT hello 6 [{101 2021-09-10 14:28:28.60466 -0700 PDT aaa 9 [] [] map[]  } {102 2021-09-10 14:28:28.60466 -0700 PDT bbb 6 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
2:      {613 2021-09-10 14:28:28.623089 -0700 PDT hello 6 [{101 2021-09-10 14:28:28.623088 -0700 PDT aaa 3 [] [] map[]  } {102 2021-09-10 14:28:28.623088 -0700 PDT bbb 3 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
3:      {613 2021-09-10 14:28:28.639698 -0700 PDT hello 9 [{101 2021-09-10 14:28:28.639697 -0700 PDT aaa 4 [] [] map[]  } {102 2021-09-10 14:28:28.639698 -0700 PDT bbb 5 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
4:      {613 2021-09-10 14:28:28.656708 -0700 PDT hello 7 [{101 2021-09-10 14:28:28.656707 -0700 PDT aaa 8 [] [] map[]  } {102 2021-09-10 14:28:28.656708 -0700 PDT bbb 0 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
5:      {613 2021-09-10 14:28:28.671936 -0700 PDT hello 5 [{101 2021-09-10 14:28:28.671935 -0700 PDT aaa 1 [] [] map[]  } {102 2021-09-10 14:28:28.671935 -0700 PDT bbb 0 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
6:      {613 2021-09-10 14:28:28.684927 -0700 PDT hello 3 [{101 2021-09-10 14:28:28.684926 -0700 PDT aaa 5 [] [] map[]  } {102 2021-09-10 14:28:28.684926 -0700 PDT bbb 8 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
7:      {613 2021-09-10 14:28:28.697836 -0700 PDT hello 3 [{101 2021-09-10 14:28:28.697835 -0700 PDT aaa 9 [] [] map[]  } {102 2021-09-10 14:28:28.697836 -0700 PDT bbb 1 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
8:      {613 2021-09-10 14:28:28.711238 -0700 PDT hello 3 [{101 2021-09-10 14:28:28.711236 -0700 PDT aaa 5 [] [] map[]  } {102 2021-09-10 14:28:28.711237 -0700 PDT bbb 3 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
9:      {613 2021-09-10 14:28:28.72443 -0700 PDT hello 1 [{101 2021-09-10 14:28:28.724429 -0700 PDT aaa 4 [] [] map[]  } {102 2021-09-10 14:28:28.72443 -0700 PDT bbb 4 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
Items with count >= 6:
0:      {613 2021-09-10 14:28:28.587935 -0700 PDT hello 7 [{101 2021-09-10 14:28:28.587934 -0700 PDT aaa 8 [] [] map[]  } {102 2021-09-10 14:28:28.587934 -0700 PDT bbb 3 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
1:      {613 2021-09-10 14:28:28.604661 -0700 PDT hello 6 [{101 2021-09-10 14:28:28.60466 -0700 PDT aaa 9 [] [] map[]  } {102 2021-09-10 14:28:28.60466 -0700 PDT bbb 6 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
2:      {613 2021-09-10 14:28:28.623089 -0700 PDT hello 6 [{101 2021-09-10 14:28:28.623088 -0700 PDT aaa 3 [] [] map[]  } {102 2021-09-10 14:28:28.623088 -0700 PDT bbb 3 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
3:      {613 2021-09-10 14:28:28.639698 -0700 PDT hello 9 [{101 2021-09-10 14:28:28.639697 -0700 PDT aaa 4 [] [] map[]  } {102 2021-09-10 14:28:28.639698 -0700 PDT bbb 5 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
4:      {613 2021-09-10 14:28:28.656708 -0700 PDT hello 7 [{101 2021-09-10 14:28:28.656707 -0700 PDT aaa 8 [] [] map[]  } {102 2021-09-10 14:28:28.656708 -0700 PDT bbb 0 [] [] map[]  foo}] [Alice Bob Cindy] map[]  }
~~~
