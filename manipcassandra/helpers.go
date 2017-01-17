package manipcassandra

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
)

// DoesKeyspaceExist checks if the configured keyspace exists
func DoesKeyspaceExist(servers []string, version int, keyspace string) (bool, error) {

	session, err := createNativeSession(servers, "", version, ExtendedTimeout)
	if err != nil {
		return false, err
	}
	defer session.Close()

	info, err := session.KeyspaceMetadata(keyspace)

	if err != nil {
		log.WithFields(logrus.Fields{
			"keyspace": keyspace,
			"error":    err,
		}).Error("unable to get keyspace metadata")

		return false, err
	}

	return len(info.Tables) > 0, nil
}

// CreateKeySpace creates a new keyspace
func CreateKeySpace(servers []string, version int, keyspace string, replicationFactor int) error {

	session, err := createNativeSession(servers, "", version, ExtendedTimeout)
	if err != nil {
		return err
	}
	defer session.Close()

	return session.Query(
		fmt.Sprintf("CREATE KEYSPACE %s WITH replication = {'class' : 'SimpleStrategy', 'replication_factor': %d}",
			keyspace,
			replicationFactor,
		)).Exec()
}

// DropKeySpace deletes the given keyspace
func DropKeySpace(servers []string, version int, keyspace string) error {

	session, err := createNativeSession(servers, "", version, ExtendedTimeout)
	if err != nil {
		return err
	}
	defer session.Close()

	return session.Query(fmt.Sprintf("DROP KEYSPACE IF EXISTS %s", keyspace)).Exec()
}

// ExecuteScript opens a new session, runs the given script in a mode and close the session.
func ExecuteScript(servers []string, version int, keyspace string, data string) error {

	session, err := createNativeSession(servers, keyspace, version, ExtendedTimeout)

	if err != nil {
		return err
	}
	defer session.Close()

	for _, statement := range strings.Split(data, ";\n") {

		if len(statement) == 0 {
			continue
		}

		if err := session.Query(statement).Exec(); err != nil {
			log.WithError(err).Error("unable to execute query. aborting script in the middle. be sure to clean up my mess.")
			return err
		}
	}

	return nil
}
