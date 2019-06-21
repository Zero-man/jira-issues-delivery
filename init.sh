#! /bin/bash
#run main
/local/main

# configure ssmtp
rm /etc/ssmtp/ssmtp.conf
cat > /etc/ssmtp/ssmtp.conf <<EOF
root=${MAIL_ROOT}
mailhub=${MAILHUB}:${PORT}
AuthUser=${MAIL_USER}
AuthPass=${MAIL_PASSWORD}
UseTLS=Yes
UseSTARTTLS=Yes
hostname=${MAIL_HOSTNAME}
FromLineOverride=YES
EOF

#email each markdown file
for file in ./markdown/*.md; do    
    cmd="echo -e \"to: $RECIPIENT\nsubject: ${file//.md}\n\" | cat - && uuencode $file ${file##*/} | ssmtp $RECIPIENT"
    eval output=\`$cmd\`
    echo $output
done