package models

import (
	"database/sql"
	"encoding/json"

	"gorm.io/datatypes"
)

type ProfileInventory struct {
	BaseModel
	ProfileID          string         `gorm:"primaryKey" json:"-"`
	Items              []ProfileItem  `json:"items" gorm:"foreignKey:ProfileInventoryID;references:ID"`
	Equipment          *ProfileItemID `json:"equipment" gorm:"foreignKey:ProfileInventoryID;references:ID"`
	Stash              *ProfileItemID `json:"stash" gorm:"foreignKey:ProfileInventoryID;references:ID"`
	SortingTable       *ProfileItemID `json:"sortingTable" gorm:"foreignKey:ProfileInventoryID;references:ID"`
	QuestRaidItems     *ProfileItemID `json:"questRaidItems" gorm:"foreignKey:ProfileInventoryID;references:ID"`
	QuestStashItems    *ProfileItemID `json:"questStashItems" gorm:"foreignKey:ProfileInventoryID;references:ID"`
	FastPanel          datatypes.JSON `json:"fastPanel" gorm:"default:'{}'"`
	HideoutAreaStashes datatypes.JSON `json:"hideoutAreaStashes" gorm:"default:'{}'"`
	FavoriteItems      datatypes.JSON `json:"favoriteItems" gorm:"default:'[]'"`
}

type ProfileItem struct {
	BaseModel
	ID                             string          `gorm:"primaryKey" json:"_id"`
	ProfileInventoryID             string          `gorm:"primaryKey" json:"-"`
	Tpl                            string          `json:"_tpl"`
	ParentID                       string          `json:"parentId,omitempty"`
	SlotID                         string          `json:"slotId,omitempty"`
	LocationX                      sql.NullInt16   `json:"-" gorm:"column:location_x"`
	LocationY                      sql.NullInt16   `json:"-" gorm:"column:location_y"`
	Rotation                       sql.NullInt16   `json:"-" gorm:"column:rotation"`
	IsSearched                     sql.NullBool    `json:"-" gorm:"column:is_searched"`
	SpawnedInSession               sql.NullBool    `json:"-" gorm:"column:spawned_in_session"`
	MaxDurability                  sql.NullFloat64 `json:"-" gorm:"column:max_durability"`
	Durability                     sql.NullFloat64 `json:"-" gorm:"column:durability"`
	FireMode                       sql.NullString  `json:"-" gorm:"column:fire_mode"`
	ScopesCurrentCalibPointIndexes datatypes.JSON  `json:"-" gorm:"column:scopes_current_calib_point_indexes"`
	ScopesSelectedModes            datatypes.JSON  `json:"-" gorm:"column:scopes_selected_modes"`
	SelectedScope                  sql.NullInt16   `json:"-" gorm:"column:selected_scope"`
}

type ProfileItemID string

// MarshalJSON for ItemID to convert to a JSON string
func (id *ProfileItemID) MarshalJSON() ([]byte, error) {
	if id == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(string(*id))
}

// UnmarshalJSON for ItemID to convert from a JSON string
func (id *ProfileItemID) UnmarshalJSON(data []byte) error {
	var str *string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	if str == nil {
		*id = ""
	} else {
		*id = ProfileItemID(*str)
	}
	return nil
}

type Location struct {
	X          int   `json:"x,omitempty"`
	Y          int   `json:"y,omitempty"`
	R          int   `json:"r,omitempty"`
	IsSearched *bool `json:"isSearched,omitempty"`
}

type Upd struct {
	SpawnedInSession bool        `json:"SpawnedInSession,omitempty"`
	Repairable       *Repairable `json:"Repairable,omitempty"`
	FireMode         *FireMode   `json:"FireMode,omitempty"`
	Sight            *Sight      `json:"Sight,omitempty"`
}

type Repairable struct {
	MaxDurability float64 `json:"MaxDurability,omitempty"`
	Durability    float64 `json:"Durability,omitempty"`
}

type FireMode struct {
	FireMode string `json:"FireMode,omitempty"`
}

type Sight struct {
	ScopesCurrentCalibPointIndexes []int `json:"ScopesCurrentCalibPointIndexes,omitempty"`
	ScopesSelectedModes            []int `json:"ScopesSelectedModes,omitempty"`
	SelectedScope                  int   `json:"SelectedScope,omitempty"`
}

func (item *ProfileItem) MarshalJSON() ([]byte, error) {
	type Alias ProfileItem
	var location *Location = nil
	var upd *Upd = nil
	var isSearched *bool = nil
	if item.IsSearched.Valid {
		isSearched = &item.IsSearched.Bool
	}
	if item.LocationX.Valid && item.LocationY.Valid && item.Rotation.Valid && item.IsSearched.Valid {
		location = &Location{
			X:          int(item.LocationX.Int16),
			Y:          int(item.LocationY.Int16),
			R:          int(item.Rotation.Int16),
			IsSearched: isSearched,
		}
	}
	var repairable *Repairable = nil
	var fireMode *FireMode = nil
	var sight *Sight = nil
	if item.MaxDurability.Valid && item.Durability.Valid {
		repairable = &Repairable{
			MaxDurability: item.MaxDurability.Float64,
			Durability:    item.Durability.Float64,
		}
	}

	if item.FireMode.Valid {
		fireMode = &FireMode{
			FireMode: item.FireMode.String,
		}

	}

	if item.SelectedScope.Valid {
		var scopesCurrentCalibPointIndexes []int
		var scopesSelectedModes []int
		err := json.Unmarshal(item.ScopesCurrentCalibPointIndexes, &scopesCurrentCalibPointIndexes)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(item.ScopesSelectedModes, &scopesSelectedModes)
		if err != nil {
			return nil, err
		}
		sight = &Sight{
			ScopesCurrentCalibPointIndexes: scopesCurrentCalibPointIndexes,
			ScopesSelectedModes:            scopesSelectedModes,
			SelectedScope:                  int(item.SelectedScope.Int16),
		}
	}

	if item.SpawnedInSession.Valid || item.MaxDurability.Valid || item.Durability.Valid || item.FireMode.Valid || item.ScopesCurrentCalibPointIndexes != nil || item.ScopesSelectedModes != nil || item.SelectedScope.Valid {
		upd = &Upd{
			SpawnedInSession: item.SpawnedInSession.Valid && item.SpawnedInSession.Bool,
			Repairable:       repairable,
			FireMode:         fireMode,
			Sight:            sight,
		}
	}
	return json.Marshal(&struct {
		*Alias
		Location *Location `json:"location,omitempty"`
		Upd      *Upd      `json:"upd,omitempty"`
	}{
		Alias:    (*Alias)(item),
		Location: location,
		Upd:      upd,
	})
}
