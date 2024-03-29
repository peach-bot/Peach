build () {
    if [[ $args == *"r"* ]]
    then
        sudo systemctl stop peach
    fi

    printf "Building project\n"

    # make directories for build and hash files
    mkdir -p build || fail
    mkdir -p scripts/hash || fail
    if [[ $args == *"f"* ]]
    then
        rm -r scripts/hash || fail
        mkdir -p scripts/hash || fail
    fi

    printf "Copying files..."
    cp launchcfg.json build/launchcfg.json

    # generate service files
    echo "s/\${version}/$version/" > ./scripts/replace.txt
    sed -f ./scripts/replace.txt ./scripts/launcher_service_template.txt > build/peach_launcher.service
    sed -f ./scripts/replace.txt ./scripts/coordinator_service_template.txt > build/peach_coordinator.service
    rm ./scripts/replace.txt


    printf "done\n"

    printf "Collecting dependencies..."
    if [[ $args == *"d"* ]]
    then
    go mod download
    fi
    printf "done\n"

    builddiscordclient
    buildcoordinator
    buildlauncher
    printf "Build complete\n"
    if [[ $args == *"i"* ]]
    then
        waittillstopped
        cp build/. /home/peach -r || fail
        cp build/peach.service /etc/systemd/system/peach.service
        sudo systemctl daemon-reload
    fi
    if [[ $args == *"r"* ]]
    then
        sudo systemctl start peach
    fi
}

hash () {
    newhash=$(find ./src/$1 -type f -print0  | xargs -0 sha1sum)
    echo $newhash > scripts/hash/$1_new.hash
    newhash=$(<scripts/hash/$1_new.hash)
    rm scripts/hash/$1_new.hash
    if [[ -f "scripts/hash/$1.hash" ]];
    then
        oldhash=$(<scripts/hash/$1.hash)
    else
        oldhash=""
    fi
    if [[ "$oldhash" == "$newhash" ]];
    then
        retval=1
    else
        retval=0
    fi
    return "$retval"
}

storehash () {
    newhash=$(find ./src/$1 -type f -print0  | xargs -0 sha1sum)
    echo $newhash > scripts/hash/$1.hash
}

waittillstopped() {
    retries=20
    stopped=false
    while [ retries > 0 ] && [ stopped == false ]
    do
        systemctl is-active --quiet service && stopped=true
        sleep 1
    done
    if [ stopped == false ]
    then
        echo "Service still running after waiting for 20 seconds!"
        fail
    fi
}

buildcoordinator() {
    printf "Building coordinator"

    #check hash
    hash "peach_coordinator"
    h=$?
    if [[ "$h" == "1" ]]; then 
        printf "\nSkipping. No changes were made.\n"
        return
    fi

    printf "\nCompiling..."
    go build -o build/coordinator-$version.exe ./src/peach_coordinator || fail
    printf "\nDone building coordinator\n"
    storehash "peach_coordinator"
}

builddiscordclient() {
    printf "Building discord client"

    #check hash
    hash "peach_discord_client"
    h=$?
    if [[ "$h" == "1" ]]; then 
        printf "\nSkipping. No changes were made.\n"
        return
    fi

    printf "\nCompiling..."
    
    echo "s/\${version}/$version/" > ./scripts/replace.txt
    sed -f ./scripts/replace.txt ./scripts/version_template.txt > src/peach_discord_client/version.go
    rm ./scripts/replace.txt

    go build -o build/discordclient-$version.exe ./src/peach_discord_client || fail
    printf "\nDone building discord client\n"
    storehash "peach_discord_client"
}

buildlauncher() {
    printf "Building launcher"

    #check hash
    hash "peach_launcher"
    h=$?
    if [[ "$h" == "1" ]]; then 
        printf "\nSkipping. No changes were made.\n"
        return
    fi

    printf "\nCompiling..."

    echo "s/\${version}/$version/" > ./scripts/replace.txt
    sed -f ./scripts/replace.txt ./scripts/version_template.txt > src/peach_launcher/version.go
    rm ./scripts/replace.txt

    go build -o build/launcher-$version.exe ./src/peach_launcher || fail
    printf "\nDone building launcher\n"
    storehash "peach_launcher"
}

fail() {
    printf "Build failed\n"
    exit
}

args=$1
if [[ $args == "-"* ]]
then
    if [[ $args == *"h"* ]]
    then
        printf "Builds the project\nUsage: ./build.sh [OPTIONS]\n\nOptions:\n    -h  prints this page\n    -f  ignore cache and build all modules\n    -i  installs built project\n    -d  installs dependencies\n    -r  restarts the system service\n"
        exit
    fi
fi

version=$(./scripts/version.sh)

build