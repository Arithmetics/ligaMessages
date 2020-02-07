#!/bin/bash
if [ "$1" = "--help" ] ; then
  printf "\nWelcome to ./The localhost database restorer
  Run with no flags to download and restore from newest dropbox files
    Optional flags: 
        -d <directory>          Use a different containing dir for .tar files and mongo folder
  \n"
  exit 0
fi

autoMigrationFolder="$PWD"
chosenDirectory=""

while getopts ":d" opt; do
  case $opt in
    d)
      chosenDirectory=${OPTARG}
    ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
    :)
      echo "Option -$OPTARG requires an argument." >&2
      exit 1
      ;;
  esac
done

function main {
  copyBackupFiles $chosenDirectory
  dropAndRestorePGDatabase "claims"
  dropAndRestorePGDatebase "medicare"
  dropAndRestorePGDatabase "reports"
  dropAndRestoreMongoDatabase
  cleanUpTempFiles

  echo "Complete! Your local databases now match most the recent production backups"
}


function dropAndRestorePGDatabase {
  database=$1
  psql -U postgres -h localhost -p 5432 -d ${database} -w -f "$autoMigrationFolder/delete${database}SchemaNoCreate.sql"
  echo "Restoring PG ${database}DB from Backup"
  pg_restore -U postgres -h localhost -p 5432 -w -d claims "$autoMigrationFolder/tempBackups/${database}.tar"
  if [ $? != 0 ]; then
    echo "Error in pg_restore of ${database}db"
    exit 0
  fi
}

function dropAndRestoreMongoDatabase {
  echo "Restore MongoDB from Backup"
  mongorestore --host localhost:27017 --drop "$autoMigrationFolder/tempBackups/mongo"
}

function copyFilesFromDropbox {
  echo "Downloading most recent backups from Dropbox"
  cd ~/Dropbox\ \(6\ Degrees\ Health\)/Software\ Development
  cd Claims/DataDumps
  echo "restoring from backup: $(\ls -1dt ./*/ | head -n 1)"
  cp -a "$(\ls -1dt ./*/ | head -n 1)" "$autoMigrationFolder/tempBackups"
  sleep 1s
}

function copyFilesFromDir {
  dir=$1
  echo "Copying files from specified directory: ${dir}"
  cd ${dir}
  cp -a "${dir}"/. "$autoMigrationFolder/tempBackups"
  sleep 1s
}

function copyBackupFiles {
  if [[ $1 == "" ]] 
  then
    copyFilesFromDropbox
  else
    copyFilesFromDir
  fi
}

function cleanUpTempFiles {
  rm -rf $autoMigrationFolder/tempBackups/*
}

main