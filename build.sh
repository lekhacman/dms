app=dms

echo "Building" $app
go build -o dist/$app

chmod +x dist/$app

buildah bud -t $app .
