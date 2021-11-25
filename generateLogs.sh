#!/bin/bash
ipsum="Aliquip labore anim et amet ipsum id proident nisi. Fugiat anim id tempor elit aliquip ex incididunt aliquip mollit minim incididunt velit tempor nulla. Adipisicing duis non aliqua nisi eu. Voluptate elit nisi consectetur eiusmod voluptate ex ullamco dolor ad consequat nulla. Ad reprehenderit minim sint ex quis mollit reprehenderit laborum eu ut veniam nisi laboris.

Commodo nulla quis ad enim voluptate eiusmod ullamco voluptate ex ex in dolor occaecat eiusmod. Anim dolore officia mollit in dolore enim in et. Sunt consequat dolore deserunt aliquip. Incididunt exercitation adipisicing sint esse. Occaecat dolore elit Lorem excepteur officia aute nostrud.

Irure deserunt ea laborum deserunt quis mollit tempor do. Reprehenderit tempor culpa pariatur proident veniam duis eiusmod laborum exercitation occaecat aute irure labore. Do consequat reprehenderit sint ex esse incididunt exercitation laborum. Lorem excepteur aliqua esse eu qui eu ad officia. Dolore quis pariatur eu labore in labore cillum laborum qui proident. Tempor mollit amet cupidatat commodo velit culpa cupidatat culpa labore deserunt exercitation proident tempor.

Consectetur sit id pariatur ex labore non do labore cupidatat adipisicing cupidatat. Ad nulla laborum cupidatat nisi pariatur qui excepteur voluptate do qui anim amet quis cupidatat. Anim excepteur duis culpa aute laboris sunt elit nostrud id.

Magna ipsum do proident dolor mollit non qui cillum velit. Anim mollit eiusmod ex deserunt dolore. Deserunt culpa mollit deserunt sit eiusmod nisi. Aliquip et nostrud ex qui proident officia consequat pariatur ad aliqua eiusmod nulla. Eiusmod aliqua anim aliqua sunt anim aute. Ullamco ad nulla incididunt nulla.

Esse adipisicing cillum nisi anim id commodo fugiat. Dolore amet ad id adipisicing aliquip mollit deserunt mollit deserunt occaecat deserunt. Ex eu pariatur enim exercitation laboris dolor id sint.

Nulla occaecat qui ipsum ex excepteur officia duis exercitation. Irure aliqua tempor irure quis laboris labore elit non Lorem. Excepteur ea proident anim elit mollit. Velit pariatur occaecat et velit elit veniam aliqua nulla cillum aliqua proident laborum esse qui.

Tempor veniam elit ipsum commodo esse. Sit sit eu ad nostrud ipsum sunt sint. Laborum nulla do consectetur quis sunt eiusmod labore irure amet occaecat occaecat.

Dolore labore ex duis dolor non. Cillum consequat labore excepteur incididunt laboris id et id aute. Labore consequat occaecat ea mollit do sint est labore irure officia do qui. Pariatur occaecat aliquip nulla sunt enim cupidatat duis do Lorem. Nulla Lorem adipisicing enim proident laboris. Officia adipisicing occaecat irure exercitation velit dolor id. Non laborum proident irure nulla dolore culpa commodo ut quis.

Aute in fugiat fugiat excepteur non sint dolor ut id veniam nostrud. Veniam irure qui ut consectetur ex elit. Ex cupidatat veniam tempor adipisicing occaecat duis Lorem officia. Duis sunt sit enim laboris pariatur minim laborum veniam qui occaecat aliquip."

# This config generates about 717Mb of logfiles
# with 500 files 1.5Mb per logfile
content=""

amntFiles=$1
length=$2

outputFolder=$3

for i in $(seq 1 $length); do
    content="$content$ipsum"
    echo -ne "\rAdding Content $i/$length"
done

mkdir -p $outputFolder/
mkdir -p $outputFolder/nested/
echo "Generating $amntFiles files..."
for i in $(seq 1 $amntFiles); do
    echo "$content" >$outputFolder/log.$i.txt &
    echo -ne "\rCreating file $i/$amntFiles"
done

for i in $(seq 1 $amntFiles); do
    echo "$content" >$outputFolder/nested/log.$i.txt &
    echo -ne "\rCreating nested file $i/$amntFiles"
done

echo -e "\nDone!"
echo "Generated $(du -sh $outputFolder/ | cut -f1) of logs"
