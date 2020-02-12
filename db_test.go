package main

import (
	"os"
	"testing"
)

func Test_DB_init(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		db := &DB{}
		err := db.init()
		if err != nil || db.conn == nil {
			t.Errorf("%s", err.Error())
		}
		defer db.conn.Close()

		_, err = os.Stat("./dungeonbot.go")
		if err != nil {
			t.Errorf("%s", err.Error())
		}

		_, err = db.conn.Exec("INSERT OR REPLACE INTO pcs (nick, campaignName, charName, notes) VALUES(?, ?, ?, ?);", "foobat", "testCampaign", "testPlayer", "some notes")
		if err != nil {
			t.Errorf("%s", err.Error())
		}

		row := Row{}
		tmprow := db.conn.QueryRow("SELECT * FROM pcs WHERE campaignName='testCampaign'")
		err = tmprow.Scan(&row.nick, &row.campaign, &row.char, &row.notes)
		if err != nil {
			t.Errorf("%s", err.Error())
		}
		if row.nick != "foobat" {
			t.Errorf("Did not retrieve nick name")
		}
		if row.campaign != "testCampaign" {
			t.Errorf("Did not retrieve campaign name")
		}
		if row.char != "testPlayer" {
			t.Errorf("Did not retrieve player name")
		}
		if row.notes != "some notes" {
			t.Errorf("Did not retrieve notes")
		}
	})
}