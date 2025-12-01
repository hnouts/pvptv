-- Seed WoW classes (excluding Plunderstorm which is being removed)
INSERT INTO classes (game_slug, class_slug, class_name) VALUES
  ('wow', 'death_knight', 'Death Knight'),
  ('wow', 'demon_hunter', 'Demon Hunter'),
  ('wow', 'druid', 'Druid'),
  ('wow', 'evoker', 'Evoker'),
  ('wow', 'hunter', 'Hunter'),
  ('wow', 'mage', 'Mage'),
  ('wow', 'monk', 'Monk'),
  ('wow', 'paladin', 'Paladin'),
  ('wow', 'priest', 'Priest'),
  ('wow', 'rogue', 'Rogue'),
  ('wow', 'shaman', 'Shaman'),
  ('wow', 'warlock', 'Warlock'),
  ('wow', 'warrior', 'Warrior')
ON CONFLICT (game_slug, class_slug) DO NOTHING;

