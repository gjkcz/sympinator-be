--------------------------------------------------------------------------
-- GET collections by users
--------------------------------------------------------------------------

SELECT u.name as Person_Name, u.id as Person, c.id as Collection, c.collection_name as Collection_Name FROM 'users' u
	INNER JOIN 'collections_by_users' r ON r.users_id = u.id
	INNER JOIN 'collections' c ON r.collections_id = c.id
ORDER BY g.sort_order
--------------------------------------------------------------------------
-- GET highest hierarchical position by user id
--------------------------------------------------------------------------

SELECT u.name as Person_Name, u.id as Person, c.id as Collection, c.collection_name as Collection_Name, c.hierarchy_pos as Status FROM users u
	INNER JOIN collections_by_users r ON r.users_id = u.id
	INNER JOIN collections c ON r.collections_id = c.id
ORDER BY c.hierarchy_pos LIMIT 1;
WHERE u.id=?

--------------------------------------------------------------------------
-- GET people not contained in any group which is higher up than current user's
--------------------------------------------------------------------------
-- TODO: it is possible this may be made significantly faster by pre-computing
-- the users-to-collections table

SELECT u.id as Person, u.name as Person_name FROM 'users' u
WHERE u.id NOT IN (
SELECT u.id FROM 'users' u
	INNER JOIN 'collections_by_users' r ON r.users_id = u.id
	INNER JOIN 'collections' c ON r.collections_id = c.id
	WHERE c.hierarchy_pos <= (
		SELECT MIN(c.hierarchy_pos) FROM 'users' u
		INNER JOIN 'collections_by_users' r ON r.users_id = u.id
		INNER JOIN 'collections' c ON r.collections_id = c.id
		WHERE u.id = ?
	)
);

--------------------------------------------------------------------------
-- GET group's permissions
--------------------------------------------------------------------------

SELECT c.id as Collection, c.collection_name as Collection_Name,
		p.id as Privilege, p.name as Privilege_name FROM collections c
	INNER JOIN permissions_by_collections r ON r.collection_id = c.id
	INNER JOIN user_permissions ON r.collection_id = p.id
ORDER BY g.hierarchy_pos;
WHERE u.id=?
--------------------------------------------------------------------------
-- GET permissions by user id -- implemented
--------------------------------------------------------------------------

SELECT DISTINCT u.name as Person_Name, u.id as Person, p.id as Permission, p.name as Permission_name FROM 'users' u
	INNER JOIN 'collections_by_users' r1 ON r1.users_id = u.id
	INNER JOIN 'permissions_by_collections' r2 ON r1.collections_id = r2.collections_id
	INNER JOIN 'user_permissions' p ON r2.perm_id = p.id
WHERE u.id=?
