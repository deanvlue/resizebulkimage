 archivo="$1"

 while read -r line
    do
         name="$line"
         #echo "https://cms-rewards.starbucks.mx/cms/cards/art/$name"
         curl -O "https://cms-rewards.starbucks.mx/cms/cards/art/$name"
    done < "$archivo"

