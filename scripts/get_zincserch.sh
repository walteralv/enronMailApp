
if [[ ! -d "Zinc/" ]]; then

    if [[ ! -f "zinc_0.3.5_Linux_x86_64.tar.gz" ]]; then
        echo "Downloading zincserch from https://github.com/zinclabs/zinc/releases/download/v0.3.5/zinc_0.3.5_Linux_x86_64.tar.gz ..."
        wget https://github.com/zinclabs/zinc/releases/download/v0.3.5/zinc_0.3.5_Linux_x86_64.tar.gz 
    fi
    mkdir Zinc
    mkdir data
    echo "Extracting the zincserch..."
    tar -C Zinc/ -xvzf zinc_0.3.5_Linux_x86_64.tar.gz
    echo "Extracted successfully!"

else
    echo "Directory 'Zinc' already exists"
fi


