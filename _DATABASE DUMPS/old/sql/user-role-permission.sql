SELECT "user".*, roles.* 
FROM "user",
LATERAL (
	SELECT array_to_json(array_agg(role_agg)) AS roles
	FROM (
		SELECT role.id, role.name, inner1.permissions 
		FROM role,
		LATERAL (
			SELECT array_to_json(array_agg(permission_agg)) AS permissions
			FROM ( 
				SELECT * FROM permission
				WHERE id = ANY (role.permission_ids)
			) permission_agg
		) inner1
		WHERE role.id = ANY ("user".role_ids)
	) role_agg
) roles
WHERE id=2;

SELECT row_to_json(t)
	FROM (
	SELECT "user".*, roles.* 
	FROM "user",
	LATERAL (
		SELECT array_to_json(array_agg(role_agg)) AS roles
		FROM (
			SELECT role.id, role.name, inner1.permissions 
			FROM role,
			LATERAL (
				SELECT array_to_json(array_agg(permission_agg)) AS permissions
				FROM ( 
					SELECT * FROM permission
					WHERE id = ANY (role.permission_ids)
				) permission_agg
			) inner1
			WHERE role.id = ANY ("user".role_ids)
		) role_agg
	) roles
	WHERE id=2
)t