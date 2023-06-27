cd ../../charts
for d in *
do
  echo "Testing $d "
  (cd "$d" && helm lint )
  if [ $? -ne 0 ]; then
    echo "Error"
    exit 1
  fi
done

