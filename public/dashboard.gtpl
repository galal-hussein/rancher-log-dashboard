<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Log Dashboard - Rancher Infra stack</title>

  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css">
  {{ if eq (.TemplateType) "container" }}
  <link rel="stylesheet" href="../static/css/dashboard.css">
  {{ else }}
  <link rel="stylesheet" href="./static/css/dashboard.css">
  {{ end }}
</head>

<body>
  <nav class="navbar navbar-fixed-top">
    <div class="container">
      <div class="navbar-header">
        <a class="navbar-brand" href="/">Log Dashboard<span class="coligo"></span></a>
      </div>
    </div>
  </nav>

  <div class="container">
    <div id="app">
      <div class="row">
        <div class="col-lg-3">
          <div class="well">
            <h3>Infra Stacks</h3>
            <div class="stacks">
              <a href="/healthcheck" class="btn btn-primary btn-block" role="button">Healthcheck</a>
              <a href="/ipsec" class="btn btn-primary btn-block" role="button">Ipsec</a>
              <a href="/network" class="btn btn-primary btn-block" role="button">Network Services</a>
              <a href="/scheduler" class="btn btn-primary btn-block" role="button">Scheduler</a>
            </div>
          </div>
        </div>
        <div class="col-lg-9">
            {{ if eq (.TemplateType) "service" }}
          <div class="well">
          <h3>Containers</h2>
              {{ range $id, $name := .Containers }}
              <a href="/healthcheck/{{ $id }}" class="btn btn-primary btn-block" role="button">{{ $name.Name }}</a>
              {{end}}
          </div>
            {{ else if eq (.TemplateType) "container" }}
              <code>
              {{ range $i, $log := .C.Logs }}
              {{ $log }}<br />
              {{end}}
              </code>
            {{end}}
      </div>
      </div>
    </div>
  </div>

</body>

</html>
