<?php
    $titulo = "Gráfico de Pizza - Coberturas de Pizza";
    $dados = [
        ['Cogumelos', 3],
        ['Cebolas', 1],
        ['Azeitonas', 1],
        ['Pimentões', 2],
        ['Salsicha', 2]
    ];
?>

<!DOCTYPE html>
<html lang="en">

<head>
    <script>
        function drawChart() {

            var data = new google.visualization.DataTable();
            data.addColumn('string', 'Cobertura');
            data.addColumn('number', 'Fatias');
            data.addRows(<?php echo json_encode($dados); ?>);

            var options = {
                title: '<?php echo $titulo; ?>',
                width: 400,
                height: 300,
                is3D: true
            };

            var chart = new google.visualization.BarChart(document.getElementById('chart_div'));
            chart.draw(data, options);
        }
    </script>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>

<body>
    <script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
    <script type="text/javascript">
        google.charts.load('current', { 'packages': ['corechart'] });
        google.charts.setOnLoadCallback(drawChart);
    </script>
    <div id="chart_div"></div>
</body>

</html>