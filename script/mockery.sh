# Run this command from root project directory
# sh scripts/mockery.sh

# Create mock for all interfaces
rm -rf ./mocks
mockery --all --with-expecter --case underscore
