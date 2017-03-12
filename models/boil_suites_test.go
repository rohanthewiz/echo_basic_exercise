package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Languages", testLanguages)
	t.Run("Jets", testJets)
	t.Run("Pilots", testPilots)
}

func TestDelete(t *testing.T) {
	t.Run("Languages", testLanguagesDelete)
	t.Run("Jets", testJetsDelete)
	t.Run("Pilots", testPilotsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Languages", testLanguagesQueryDeleteAll)
	t.Run("Jets", testJetsQueryDeleteAll)
	t.Run("Pilots", testPilotsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Languages", testLanguagesSliceDeleteAll)
	t.Run("Jets", testJetsSliceDeleteAll)
	t.Run("Pilots", testPilotsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Languages", testLanguagesExists)
	t.Run("Jets", testJetsExists)
	t.Run("Pilots", testPilotsExists)
}

func TestFind(t *testing.T) {
	t.Run("Languages", testLanguagesFind)
	t.Run("Jets", testJetsFind)
	t.Run("Pilots", testPilotsFind)
}

func TestBind(t *testing.T) {
	t.Run("Languages", testLanguagesBind)
	t.Run("Jets", testJetsBind)
	t.Run("Pilots", testPilotsBind)
}

func TestOne(t *testing.T) {
	t.Run("Languages", testLanguagesOne)
	t.Run("Jets", testJetsOne)
	t.Run("Pilots", testPilotsOne)
}

func TestAll(t *testing.T) {
	t.Run("Languages", testLanguagesAll)
	t.Run("Jets", testJetsAll)
	t.Run("Pilots", testPilotsAll)
}

func TestCount(t *testing.T) {
	t.Run("Languages", testLanguagesCount)
	t.Run("Jets", testJetsCount)
	t.Run("Pilots", testPilotsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Languages", testLanguagesHooks)
	t.Run("Jets", testJetsHooks)
	t.Run("Pilots", testPilotsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Languages", testLanguagesInsert)
	t.Run("Languages", testLanguagesInsertWhitelist)
	t.Run("Jets", testJetsInsert)
	t.Run("Jets", testJetsInsertWhitelist)
	t.Run("Pilots", testPilotsInsert)
	t.Run("Pilots", testPilotsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("JetToPilotUsingPilot", testJetToOnePilotUsingPilot)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("LanguageToPilots", testLanguageToManyPilots)
	t.Run("PilotToJets", testPilotToManyJets)
	t.Run("PilotToLanguages", testPilotToManyLanguages)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("JetToPilotUsingPilot", testJetToOneSetOpPilotUsingPilot)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("LanguageToPilots", testLanguageToManyAddOpPilots)
	t.Run("PilotToJets", testPilotToManyAddOpJets)
	t.Run("PilotToLanguages", testPilotToManyAddOpLanguages)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Languages", testLanguagesReload)
	t.Run("Jets", testJetsReload)
	t.Run("Pilots", testPilotsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Languages", testLanguagesReloadAll)
	t.Run("Jets", testJetsReloadAll)
	t.Run("Pilots", testPilotsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Languages", testLanguagesSelect)
	t.Run("Jets", testJetsSelect)
	t.Run("Pilots", testPilotsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Languages", testLanguagesUpdate)
	t.Run("Jets", testJetsUpdate)
	t.Run("Pilots", testPilotsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Languages", testLanguagesSliceUpdateAll)
	t.Run("Jets", testJetsSliceUpdateAll)
	t.Run("Pilots", testPilotsSliceUpdateAll)
}

func TestUpsert(t *testing.T) {
	t.Run("Languages", testLanguagesUpsert)
	t.Run("Jets", testJetsUpsert)
	t.Run("Pilots", testPilotsUpsert)
}
