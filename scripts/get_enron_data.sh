# script to get the enron mails data

if [[ ! -d "db/enron_mail_20110402" ]]; then
    
    if [[ ! -f "enron_mail_20110402.tgz" ]]; then
        echo "Downloading the data from http://download.srv.cs.cmu.edu/\~enron/enron_mail_20110402.tgz ..."
        curl -O http://download.srv.cs.cmu.edu/\~enron/enron_mail_20110402.tgz 2>&1    
    fi

    echo "Extracting the data..."
    tar -C db/ -xzf enron_mail_20110402.tgz
    echo "Extracted successfully!"
else
    echo "Directory 'enron_mail_20110402' already exists"
fi
