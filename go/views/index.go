package views

// Index html content
const Index = `
<!DOCTYPE html>
<html>
<head>
    <title>{{.appname}}</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0">
</head>
<body style="height:100%;width:100%;">
    <div id="app"></div>
</body>
<script>{{.globalconfig}}</script>
<script src="/static/js/vendor.min.js{{.unix}}"></script>
<script src="/static/js/app.min.js{{.unix}}"></script>
</html>
`
