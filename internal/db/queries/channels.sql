-- name: ListPublishedChannels :many
SELECT 
  id,
  twitch_login,
  display_name,
  is_published,
  socials,
  created_at,
  updated_at
FROM channels
WHERE is_published = true
ORDER BY sort_weight DESC, display_name ASC;

-- name: GetChannelBySlug :one
SELECT 
  id,
  twitch_login,
  display_name,
  is_published,
  socials,
  created_at,
  updated_at
FROM channels
WHERE twitch_login = $1 AND is_published = true;

-- name: ListChannelsByClass :many
-- This will be used once channel_specs are populated
SELECT DISTINCT
  c.id,
  c.twitch_login,
  c.display_name,
  c.is_published,
  c.socials,
  c.created_at,
  c.updated_at
FROM channels c
INNER JOIN channel_specs cs ON c.id = cs.channel_id
INNER JOIN specs s ON cs.spec_id = s.id
INNER JOIN classes cl ON s.class_id = cl.id
WHERE cl.game_slug = 'wow' 
  AND cl.class_slug = $1
  AND c.is_published = true
ORDER BY c.sort_weight DESC, c.display_name ASC;

-- name: GetClassBySlug :one
SELECT id, game_slug, class_slug, class_name
FROM classes
WHERE game_slug = 'wow' AND class_slug = $1;

-- name: ListAllClasses :many
SELECT id, game_slug, class_slug, class_name
FROM classes
WHERE game_slug = 'wow'
ORDER BY class_name ASC;
