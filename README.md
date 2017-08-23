

# Dependency
1. Need install "go get -u github.com/crewjam/go-xmlsec"
    This package uses cgo to wrap libxmlsec. As such, you'll need libxmlsec headers and a C compiler to make it work.
        On linux, this might look like:
            $ apt-get install libxml2-dev libxmlsec1-dev
            $ go get github.com/crewjam/go-xmlsec
        On Mac with homebrew, this might look like:
            $ brew update
            $ brew install libxmlsec1 libxml2 libxslt
           	$ echo 'export PATH="/usr/local/opt/libxslt/bin:$PATH"' >> ~/.bash_profile
           	$ echo 'export PATH="/usr/local/opt/libxml2/bin:$PATH"' >> ~/.bash_profile
           	$ echo 'export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/opt/libxml2/lib/pkgconfig:/usr/local/opt/libxslt/lib/pkgconfig' >> ~/.bash_profile
           	$ source ~/.bash_profile
            $ go get github.com/crewjam/go-xmlsec
            
# Enable Travis

./kunlun.sh -s metadata -c taishan_metadata
