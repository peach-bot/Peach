./scripts/build.sh $1

function on_exit()
{
    echo "Exiting..."
}

trap "on_exit" 2
echo "Running launcher"

version=$(git describe --tags)
branch=$(git rev-parse --abbrev-ref HEAD)
if [[ "$branch" == "master" ]]; then
    branch=""
else
    branch="-$branch"
fi
version=${version%-*-*}
version="$version$branch"

cd build
./launcher-$version.exe