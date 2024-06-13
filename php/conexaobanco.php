<?php

	session_start();

	$host="127.0.0.1";
	$port=3306;
	$socket="";
	$user="vittas";
	$password="RPYYtQse[T~3s1(";
	$dbname="bancovitor";

	$conn = new mysqli($host, $user, $password, $dbname, $port, $socket);
		
	if ($conn->connect_error) {
		die("Connection failed: " . $conn->connect_error);
	  }

	  
?>


