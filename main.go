package main

import "fmt";

func main()  {
	variable() ;
}

// func helloWorld()  {
// 	fmt.Print("Hello Worlds");
// }

func variable()  {
	// Deklarasi Variabel Beserta Tipe Data
	var lastName string
	var firstName string = "Ardisca";

	lastName ="Evanandy";


	fmt.Printf("halo %s %s!\n", firstName, lastName);

	//Deklarasi Variabel Tanpa Tipe Data
	var firstName1 string= "Lord";
	lastName1 :="Voldemor";

	fmt.Printf("Hello %s %s!\n", firstName1, lastName1);

	// Deklarasi Multi Variabel
	var first, sec, third string;
	
	first, sec, third ="first", "sec", "third";
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"


	fmt.Printf("Hello %s %s %s\n",first, sec,third);
	fmt.Printf("%s number %d this is Friday? %t do u have $%.1f\n", say,one, isFriday, twoPointTwo );
	
	// Variabel Underscore _
	_ = "john";

	// Deklarasi Variabel Menggunakan Keyword new
	name := new(string);
	fmt.Println(name);
	fmt.Println(*name);
}