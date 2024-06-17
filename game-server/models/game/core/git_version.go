package core

import (
	"bytes"
	"game-server/helpers"
	"time"
)

type GitVersion struct {
	CommitHash    string
	CommitDate    time.Time
	CommitSubject string
	CommitBranch  string
}

func DeserializeGitVersion(buffer *bytes.Buffer) (GitVersion, error) {
	var gv GitVersion
	var err error

	if err = helpers.ReadString(buffer, &gv.CommitHash); err != nil {
		return gv, err
	}
	if err = helpers.ReadDateTime(buffer, &gv.CommitDate); err != nil {
		return gv, err
	}
	if err = helpers.ReadString(buffer, &gv.CommitSubject); err != nil {
		return gv, err
	}
	if err = helpers.ReadString(buffer, &gv.CommitBranch); err != nil {
		return gv, err
	}
	return gv, nil
}

func (gv *GitVersion) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteString(buffer, gv.CommitHash); err != nil {
		return err
	}
	if err = helpers.WriteDateTime(buffer, gv.CommitDate); err != nil {
		return err
	}
	if err = helpers.WriteString(buffer, gv.CommitSubject); err != nil {
		return err
	}
	if err = helpers.WriteString(buffer, gv.CommitBranch); err != nil {
		return err
	}
	return nil
}
