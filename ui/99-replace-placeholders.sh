#!/bin/sh
set -e -o errexit

VARS=$(env | grep -E '^VITE_.+')
APP_DIR=/usr/share/nginx/html

echo -e "Vars to be set:\n$(echo $VARS | sed 's~ ~\n~g')"

for FILE in $(find $APP_DIR -type f \( -name "*.html" -o -name "*.css" -o -name "*.js" \)); do
  echo "Processing file: $FILE"
  TMP_FILE=$(mktemp)
  cp $FILE $TMP_FILE
  for VAR in $VARS; do
    KEY=$(echo $VAR | awk -F= '{print $1}')
    VAL=$(echo $VAR | awk -F= '{print $2}')
    echo "Replacing <$KEY> with $VAL in $FILE"
    sed -i "s~<$KEY>~$VAL~g" $TMP_FILE
  done
  mv $TMP_FILE $FILE
done

echo "Placeholders have been replaced"
