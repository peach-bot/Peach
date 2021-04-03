build () {
    if [[ $args == *"r"* ]]
    then
        sudo systemctl stop peach
    fi
    printf "Building project\n"
    mkdir -p build || fail

    printf "Copying files..."
    cp launchcfg.json build/launchcfg.json
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
        cp build/. /home/peach -r || fail
        cp peach.service /etc/systemd/system/peach.service
        sudo systemctl systemctl daemon-reload
    fi
    if [[ $args == *"r"* ]]
    then
        sudo systemctl start peach
    fi
}

hash () {
    mkdir -p scripts/build || fail
    newhash=$(find ./src/$1 -type f -print0  | xargs -0 sha1sum)
    echo $newhash > scripts/build/$1_new.hash
    newhash=$(<scripts/build/$1_new.hash)
    rm scripts/build/$1_new.hash
    if [[ -f "scripts/build/$1.hash" ]];
    then
        oldhash=$(<scripts/build/$1.hash)
    else
        oldhash=""
    fi
    if [[ "$oldhash" == "$newhash" ]];
    then
        retval=1
    else
        retval=0
        echo $newhash > scripts/build/$1.hash
    fi
    return "$retval"
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
    printf "Building client coordinator"

    #check hash
    hash "peach_client_coordinator"
    h=$?
    if [[ "$h" == "1" ]]; then 
        printf "\nSkipping. No changes were made.\n"
        return
    fi

    printf "\nCompiling..."
    go build -o build/coordinator.exe ./src/peach_client_coordinator || fail
    if [[ $args == *"i"* ]]
    then
        waittillstopped
        cp build/coordinator.exe /usr/local/bin/peach/coordinator || fail
        sudo setcap CAP_NET_BIND_SERVICE=+eip /usr/local/bin/coordinator || fail
    fi
    printf "\nDone building client coordinator\n"
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
    go build -o build/discordclient.exe ./src/peach_discord_client || fail
    if [[ $args == *"i"* ]]
    then
        waittillstopped
        cp build/discordclient.exe /usr/local/bin/peach/discordclient || fail
        sudo setcap CAP_NET_BIND_SERVICE=+eip /usr/local/bin/discordclient || fail
    fi
    printf "\nDone building discord client\n"
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
    go build -o build/launcher.exe ./src/peach_launcher || fail
    if [[ $args == *"i"* ]]
    then
        waittillstopped
        cp build/launcher /usr/local/bin/peach/launcher || fail
        sudo setcap CAP_NET_BIND_SERVICE=+eip /usr/local/bin/launcher || fail
    fi
    printf "\nDone building launcher\n"
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
        printf "Builds the project\nUsage: ./build.sh [OPTIONS]\n\nOptions:\n    -h  prints this page\n    -i  installs built project\n    -d  installs dependencies\n    -r  restarts the system service\n"
        exit
    fi
fi
build