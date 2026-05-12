<?php
if (isset($_GET['action']) && $_GET['action'] == 'logout') {
    setcookie("user", "", time() - 3600, "/");
    header("Location: login.php");
    exit();
}

if (!isset($_COOKIE['user'])) {
    header("Location: login.php");
    exit();
}

$user = $_COOKIE['user'];
?>

<!DOCTYPE html>
<html>
<head>
    <title>Site</title>
</head>
<body>

<h2>Olá, <?php echo htmlspecialchars($user); ?>!</h2>

<a href="site.php?action=logout">Sair (Logout)</a>

</body>
</html>
