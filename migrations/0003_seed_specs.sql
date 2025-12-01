-- Seed WoW class specializations
-- This uses subqueries to get class IDs dynamically

-- Death Knight
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'blood', 'Blood' FROM classes WHERE class_slug = 'death_knight' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'frost', 'Frost' FROM classes WHERE class_slug = 'death_knight' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'unholy', 'Unholy' FROM classes WHERE class_slug = 'death_knight' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Demon Hunter
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'havoc', 'Havoc' FROM classes WHERE class_slug = 'demon_hunter' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'vengeance', 'Vengeance' FROM classes WHERE class_slug = 'demon_hunter' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Druid
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'balance', 'Balance' FROM classes WHERE class_slug = 'druid' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'feral', 'Feral' FROM classes WHERE class_slug = 'druid' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'guardian', 'Guardian' FROM classes WHERE class_slug = 'druid' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'restoration', 'Restoration' FROM classes WHERE class_slug = 'druid' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Evoker
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'devastation', 'Devastation' FROM classes WHERE class_slug = 'evoker' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'preservation', 'Preservation' FROM classes WHERE class_slug = 'evoker' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'augmentation', 'Augmentation' FROM classes WHERE class_slug = 'evoker' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Hunter
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'beast_mastery', 'Beast Mastery' FROM classes WHERE class_slug = 'hunter' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'marksmanship', 'Marksmanship' FROM classes WHERE class_slug = 'hunter' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'survival', 'Survival' FROM classes WHERE class_slug = 'hunter' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Mage
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'arcane', 'Arcane' FROM classes WHERE class_slug = 'mage' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'fire', 'Fire' FROM classes WHERE class_slug = 'mage' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'frost', 'Frost' FROM classes WHERE class_slug = 'mage' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Monk
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'brewmaster', 'Brewmaster' FROM classes WHERE class_slug = 'monk' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'mistweaver', 'Mistweaver' FROM classes WHERE class_slug = 'monk' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'windwalker', 'Windwalker' FROM classes WHERE class_slug = 'monk' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Paladin
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'holy', 'Holy' FROM classes WHERE class_slug = 'paladin' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'protection', 'Protection' FROM classes WHERE class_slug = 'paladin' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'retribution', 'Retribution' FROM classes WHERE class_slug = 'paladin' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Priest
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'discipline', 'Discipline' FROM classes WHERE class_slug = 'priest' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'holy', 'Holy' FROM classes WHERE class_slug = 'priest' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'shadow', 'Shadow' FROM classes WHERE class_slug = 'priest' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Rogue
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'assassination', 'Assassination' FROM classes WHERE class_slug = 'rogue' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'outlaw', 'Outlaw' FROM classes WHERE class_slug = 'rogue' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'subtlety', 'Subtlety' FROM classes WHERE class_slug = 'rogue' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Shaman
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'elemental', 'Elemental' FROM classes WHERE class_slug = 'shaman' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'enhancement', 'Enhancement' FROM classes WHERE class_slug = 'shaman' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'restoration', 'Restoration' FROM classes WHERE class_slug = 'shaman' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Warlock
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'affliction', 'Affliction' FROM classes WHERE class_slug = 'warlock' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'demonology', 'Demonology' FROM classes WHERE class_slug = 'warlock' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'destruction', 'Destruction' FROM classes WHERE class_slug = 'warlock' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

-- Warrior
INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'arms', 'Arms' FROM classes WHERE class_slug = 'warrior' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'fury', 'Fury' FROM classes WHERE class_slug = 'warrior' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

INSERT INTO specs (class_id, spec_slug, spec_name)
SELECT id, 'protection', 'Protection' FROM classes WHERE class_slug = 'warrior' AND game_slug = 'wow'
ON CONFLICT (class_id, spec_slug) DO NOTHING;

