cd ../../charts
for d in *
do
  echo "Deploying $d to kind"
  (
    set -x
    cd "$d"
    if [ -f ".no_ci" ]; then
      echo "No CI for this chart."
    else
      kubectl create namespace $d || true
      helm dependency build || true
      helm install $d . --wait --timeout 300s --namespace $d
      helm test $d
    fi
  )
  if [ $? -ne 0 ]; then
    echo "Error during deployment"
    exit 1
  else
    echo "Success ! "
    helm delete $d  --namespace $d || true 
  fi
done

