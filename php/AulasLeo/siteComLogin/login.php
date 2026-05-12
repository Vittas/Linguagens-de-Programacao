<?php
session_start();
include 'connection.php';
include 'sanitization.php';

if ($_SERVER["REQUEST_METHOD"] == "POST") {
    $user = sanitize_input($_POST['username']);
    $pass = sanitize_input($_POST['password']);

    $stmt = $conn->prepare("SELECT password FROM users WHERE username = ?");
    $stmt->bind_param("s", $user);
    $stmt->execute();
    $result = $stmt->get_result();

    if ($result->num_rows > 0) {
        $row = $result->fetch_assoc();
        if (password_verify($pass, $row['password'])) {
            setcookie("user", $user, time() + (86400 * 30), "/"); 
            header("Location: site.php");
            exit();
        } else {
            $error = "Senha Invalida.";
        }
    } else {
        $error = "Usuário não encontrado.";
    }
    $stmt->close();
    $conn->close();
}
?>

<!DOCTYPE html>
<html>
<head>
    <title>Login</title>
</head>
<body>

<h2>Login Page</h2>

<form method="post" action="<?php echo htmlspecialchars($_SERVER["PHP_SELF"]);?>">
  Usuario: <input type="text" name="username">
  <br><br>
  Senha: <input type="password" name="password">
  <br><br>
  <input type="submit" name="submit" value="Login">
</form>

<?php
if (isset($error)) {
    echo "<p style='color:red;'>$error</p>";
}
?>

</body>
</html>
