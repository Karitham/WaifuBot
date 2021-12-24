package filestore

import (
	"github.com/Karitham/WaifuBot/internal/discord"
	"github.com/Karitham/corde"
	"github.com/fxamacker/cbor"
	"go.etcd.io/bbolt"
)

type Store struct {
	db *bbolt.DB
}

func New(path string) *Store {
	db, err := bbolt.Open(path, 0666, nil)
	if err != nil {
		return nil
	}
	return &Store{db: db}
}

func (s *Store) Close() error { return s.db.Close() }

func (s *Store) Put(userID corde.Snowflake, c discord.Character) error {
	return s.db.Update(func(t *bbolt.Tx) error {
		uid, err := cbor.Marshal(userID, cbor.EncOptions{})
		if err != nil {
			return err
		}

		b, err := t.CreateBucketIfNotExists(uid)
		if err != nil {
			return err
		}

		cid, err := cbor.Marshal(c.ID, cbor.EncOptions{})
		if err != nil {
			return err
		}

		char, err := cbor.Marshal(c, cbor.EncOptions{})
		if err != nil {
			return err
		}

		return b.Put(cid, char)
	})
}

func (s *Store) Characters(userID corde.Snowflake) ([]discord.Character, error) {
	var chars []discord.Character
	err := s.db.View(func(t *bbolt.Tx) error {
		uid, err := cbor.Marshal(userID, cbor.EncOptions{})
		if err != nil {
			return err
		}

		b := t.Bucket(uid)
		if b == nil {
			return nil
		}

		return b.ForEach(func(_, v []byte) error {
			var c discord.Character
			if err := cbor.Unmarshal(v, &c); err != nil {
				return err
			}
			chars = append(chars, c)

			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	return chars, nil
}
