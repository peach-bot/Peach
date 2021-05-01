if [[ "$TRAVIS_TAG" != "" ]]; then
    version="$TRAVIS_TAG"
else
    version=$(git describe --tags)
    branch=$(git rev-parse --abbrev-ref HEAD)
    if [[ "$branch" == "master" ]]; then
        branch=""
    else
        branch="-$branch"
    fi
    version=${version%-*-*}
    version="$version$branch"
fi

echo $version