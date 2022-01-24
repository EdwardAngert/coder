// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getAPIKeyByID = `-- name: GetAPIKeyByID :one
SELECT
  id, hashed_secret, user_id, application, name, last_used, expires_at, created_at, updated_at, login_type, oidc_access_token, oidc_refresh_token, oidc_id_token, oidc_expiry, devurl_token
FROM
  api_keys
WHERE
  id = $1
LIMIT
  1
`

func (q *sqlQuerier) GetAPIKeyByID(ctx context.Context, id string) (APIKey, error) {
	row := q.db.QueryRowContext(ctx, getAPIKeyByID, id)
	var i APIKey
	err := row.Scan(
		&i.ID,
		&i.HashedSecret,
		&i.UserID,
		&i.Application,
		&i.Name,
		&i.LastUsed,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LoginType,
		&i.OIDCAccessToken,
		&i.OIDCRefreshToken,
		&i.OIDCIDToken,
		&i.OIDCExpiry,
		&i.DevurlToken,
	)
	return i, err
}

const getOrganizationByName = `-- name: GetOrganizationByName :one
SELECT
  id, name, description, created_at, updated_at, "default", auto_off_threshold, cpu_provisioning_rate, memory_provisioning_rate, workspace_auto_off
FROM
  organizations
WHERE
  name = $1
LIMIT
  1
`

func (q *sqlQuerier) GetOrganizationByName(ctx context.Context, name string) (Organization, error) {
	row := q.db.QueryRowContext(ctx, getOrganizationByName, name)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Default,
		&i.AutoOffThreshold,
		&i.CpuProvisioningRate,
		&i.MemoryProvisioningRate,
		&i.WorkspaceAutoOff,
	)
	return i, err
}

const getOrganizationMemberByUserID = `-- name: GetOrganizationMemberByUserID :one
SELECT
  organization_id, user_id, created_at, updated_at, roles
FROM
  organization_members
WHERE
  organization_id = $1
  AND user_id = $2
LIMIT
  1
`

type GetOrganizationMemberByUserIDParams struct {
	OrganizationID string `db:"organization_id" json:"organization_id"`
	UserID         string `db:"user_id" json:"user_id"`
}

func (q *sqlQuerier) GetOrganizationMemberByUserID(ctx context.Context, arg GetOrganizationMemberByUserIDParams) (OrganizationMember, error) {
	row := q.db.QueryRowContext(ctx, getOrganizationMemberByUserID, arg.OrganizationID, arg.UserID)
	var i OrganizationMember
	err := row.Scan(
		&i.OrganizationID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.Roles),
	)
	return i, err
}

const getOrganizationsByUserID = `-- name: GetOrganizationsByUserID :many
SELECT
  id, name, description, created_at, updated_at, "default", auto_off_threshold, cpu_provisioning_rate, memory_provisioning_rate, workspace_auto_off
FROM
  organizations
WHERE
  id = (
    SELECT
      organization_id
    FROM
      organization_members
    WHERE
      user_id = $1
  )
`

func (q *sqlQuerier) GetOrganizationsByUserID(ctx context.Context, userID string) ([]Organization, error) {
	rows, err := q.db.QueryContext(ctx, getOrganizationsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Organization
	for rows.Next() {
		var i Organization
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Default,
			&i.AutoOffThreshold,
			&i.CpuProvisioningRate,
			&i.MemoryProvisioningRate,
			&i.WorkspaceAutoOff,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectByOrganizationAndName = `-- name: GetProjectByOrganizationAndName :one
SELECT
  id, created_at, updated_at, organization_id, name, provisioner, active_version_id
FROM
  project
WHERE
  organization_id = $1
  AND name = $2
LIMIT
  1
`

type GetProjectByOrganizationAndNameParams struct {
	OrganizationID string `db:"organization_id" json:"organization_id"`
	Name           string `db:"name" json:"name"`
}

func (q *sqlQuerier) GetProjectByOrganizationAndName(ctx context.Context, arg GetProjectByOrganizationAndNameParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, getProjectByOrganizationAndName, arg.OrganizationID, arg.Name)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.OrganizationID,
		&i.Name,
		&i.Provisioner,
		&i.ActiveVersionID,
	)
	return i, err
}

const getProjectHistoryByProjectID = `-- name: GetProjectHistoryByProjectID :many
SELECT
  id, project_id, created_at, updated_at, name, description, storage_method, storage_source, import_job_id
FROM
  project_history
WHERE
  project_id = $1
`

func (q *sqlQuerier) GetProjectHistoryByProjectID(ctx context.Context, projectID uuid.UUID) ([]ProjectHistory, error) {
	rows, err := q.db.QueryContext(ctx, getProjectHistoryByProjectID, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProjectHistory
	for rows.Next() {
		var i ProjectHistory
		if err := rows.Scan(
			&i.ID,
			&i.ProjectID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
			&i.StorageMethod,
			&i.StorageSource,
			&i.ImportJobID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProjectsByOrganizationIDs = `-- name: GetProjectsByOrganizationIDs :many
SELECT
  id, created_at, updated_at, organization_id, name, provisioner, active_version_id
FROM
  project
WHERE
  organization_id = ANY($1 :: text [ ])
`

func (q *sqlQuerier) GetProjectsByOrganizationIDs(ctx context.Context, ids []string) ([]Project, error) {
	rows, err := q.db.QueryContext(ctx, getProjectsByOrganizationIDs, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Project
	for rows.Next() {
		var i Project
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.OrganizationID,
			&i.Name,
			&i.Provisioner,
			&i.ActiveVersionID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmailOrUsername = `-- name: GetUserByEmailOrUsername :one
SELECT
  id, email, name, revoked, login_type, hashed_password, created_at, updated_at, temporary_password, avatar_hash, ssh_key_regenerated_at, username, dotfiles_git_uri, roles, status, relatime, gpg_key_regenerated_at, _decomissioned, shell
FROM
  users
WHERE
  username = $1
  OR email = $2
LIMIT
  1
`

type GetUserByEmailOrUsernameParams struct {
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
}

func (q *sqlQuerier) GetUserByEmailOrUsername(ctx context.Context, arg GetUserByEmailOrUsernameParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmailOrUsername, arg.Username, arg.Email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Name,
		&i.Revoked,
		&i.LoginType,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TemporaryPassword,
		&i.AvatarHash,
		&i.SshKeyRegeneratedAt,
		&i.Username,
		&i.DotfilesGitUri,
		pq.Array(&i.Roles),
		&i.Status,
		&i.Relatime,
		&i.GpgKeyRegeneratedAt,
		&i.Decomissioned,
		&i.Shell,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT
  id, email, name, revoked, login_type, hashed_password, created_at, updated_at, temporary_password, avatar_hash, ssh_key_regenerated_at, username, dotfiles_git_uri, roles, status, relatime, gpg_key_regenerated_at, _decomissioned, shell
FROM
  users
WHERE
  id = $1
LIMIT
  1
`

func (q *sqlQuerier) GetUserByID(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Name,
		&i.Revoked,
		&i.LoginType,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TemporaryPassword,
		&i.AvatarHash,
		&i.SshKeyRegeneratedAt,
		&i.Username,
		&i.DotfilesGitUri,
		pq.Array(&i.Roles),
		&i.Status,
		&i.Relatime,
		&i.GpgKeyRegeneratedAt,
		&i.Decomissioned,
		&i.Shell,
	)
	return i, err
}

const getUserCount = `-- name: GetUserCount :one
SELECT
  COUNT(*)
FROM
  users
`

func (q *sqlQuerier) GetUserCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getUserCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const insertAPIKey = `-- name: InsertAPIKey :one
INSERT INTO
  api_keys (
    id,
    hashed_secret,
    user_id,
    application,
    name,
    last_used,
    expires_at,
    created_at,
    updated_at,
    login_type,
    oidc_access_token,
    oidc_refresh_token,
    oidc_id_token,
    oidc_expiry,
    devurl_token
  )
VALUES
  (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15
  ) RETURNING id, hashed_secret, user_id, application, name, last_used, expires_at, created_at, updated_at, login_type, oidc_access_token, oidc_refresh_token, oidc_id_token, oidc_expiry, devurl_token
`

type InsertAPIKeyParams struct {
	ID               string    `db:"id" json:"id"`
	HashedSecret     []byte    `db:"hashed_secret" json:"hashed_secret"`
	UserID           string    `db:"user_id" json:"user_id"`
	Application      bool      `db:"application" json:"application"`
	Name             string    `db:"name" json:"name"`
	LastUsed         time.Time `db:"last_used" json:"last_used"`
	ExpiresAt        time.Time `db:"expires_at" json:"expires_at"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time `db:"updated_at" json:"updated_at"`
	LoginType        LoginType `db:"login_type" json:"login_type"`
	OIDCAccessToken  string    `db:"oidc_access_token" json:"oidc_access_token"`
	OIDCRefreshToken string    `db:"oidc_refresh_token" json:"oidc_refresh_token"`
	OIDCIDToken      string    `db:"oidc_id_token" json:"oidc_id_token"`
	OIDCExpiry       time.Time `db:"oidc_expiry" json:"oidc_expiry"`
	DevurlToken      bool      `db:"devurl_token" json:"devurl_token"`
}

func (q *sqlQuerier) InsertAPIKey(ctx context.Context, arg InsertAPIKeyParams) (APIKey, error) {
	row := q.db.QueryRowContext(ctx, insertAPIKey,
		arg.ID,
		arg.HashedSecret,
		arg.UserID,
		arg.Application,
		arg.Name,
		arg.LastUsed,
		arg.ExpiresAt,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.LoginType,
		arg.OIDCAccessToken,
		arg.OIDCRefreshToken,
		arg.OIDCIDToken,
		arg.OIDCExpiry,
		arg.DevurlToken,
	)
	var i APIKey
	err := row.Scan(
		&i.ID,
		&i.HashedSecret,
		&i.UserID,
		&i.Application,
		&i.Name,
		&i.LastUsed,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LoginType,
		&i.OIDCAccessToken,
		&i.OIDCRefreshToken,
		&i.OIDCIDToken,
		&i.OIDCExpiry,
		&i.DevurlToken,
	)
	return i, err
}

const insertOrganization = `-- name: InsertOrganization :one
INSERT INTO
  organizations (id, name, description, created_at, updated_at)
VALUES
  ($1, $2, $3, $4, $5) RETURNING id, name, description, created_at, updated_at, "default", auto_off_threshold, cpu_provisioning_rate, memory_provisioning_rate, workspace_auto_off
`

type InsertOrganizationParams struct {
	ID          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

func (q *sqlQuerier) InsertOrganization(ctx context.Context, arg InsertOrganizationParams) (Organization, error) {
	row := q.db.QueryRowContext(ctx, insertOrganization,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Organization
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Default,
		&i.AutoOffThreshold,
		&i.CpuProvisioningRate,
		&i.MemoryProvisioningRate,
		&i.WorkspaceAutoOff,
	)
	return i, err
}

const insertOrganizationMember = `-- name: InsertOrganizationMember :one
INSERT INTO
  organization_members (
    organization_id,
    user_id,
    created_at,
    updated_at,
    roles
  )
VALUES
  ($1, $2, $3, $4, $5) RETURNING organization_id, user_id, created_at, updated_at, roles
`

type InsertOrganizationMemberParams struct {
	OrganizationID string    `db:"organization_id" json:"organization_id"`
	UserID         string    `db:"user_id" json:"user_id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	Roles          []string  `db:"roles" json:"roles"`
}

func (q *sqlQuerier) InsertOrganizationMember(ctx context.Context, arg InsertOrganizationMemberParams) (OrganizationMember, error) {
	row := q.db.QueryRowContext(ctx, insertOrganizationMember,
		arg.OrganizationID,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
		pq.Array(arg.Roles),
	)
	var i OrganizationMember
	err := row.Scan(
		&i.OrganizationID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.Roles),
	)
	return i, err
}

const insertProject = `-- name: InsertProject :one
INSERT INTO
  project (
    id,
    created_at,
    updated_at,
    organization_id,
    name,
    provisioner
  )
VALUES
  ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at, organization_id, name, provisioner, active_version_id
`

type InsertProjectParams struct {
	ID             uuid.UUID       `db:"id" json:"id"`
	CreatedAt      time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time       `db:"updated_at" json:"updated_at"`
	OrganizationID string          `db:"organization_id" json:"organization_id"`
	Name           string          `db:"name" json:"name"`
	Provisioner    ProvisionerType `db:"provisioner" json:"provisioner"`
}

func (q *sqlQuerier) InsertProject(ctx context.Context, arg InsertProjectParams) (Project, error) {
	row := q.db.QueryRowContext(ctx, insertProject,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.OrganizationID,
		arg.Name,
		arg.Provisioner,
	)
	var i Project
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.OrganizationID,
		&i.Name,
		&i.Provisioner,
		&i.ActiveVersionID,
	)
	return i, err
}

const insertProjectHistory = `-- name: InsertProjectHistory :one
INSERT INTO
  project_history (
    id,
    project_id,
    created_at,
    updated_at,
    name,
    description,
    storage_method,
    storage_source,
    import_job_id
  )
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, project_id, created_at, updated_at, name, description, storage_method, storage_source, import_job_id
`

type InsertProjectHistoryParams struct {
	ID            uuid.UUID            `db:"id" json:"id"`
	ProjectID     uuid.UUID            `db:"project_id" json:"project_id"`
	CreatedAt     time.Time            `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time            `db:"updated_at" json:"updated_at"`
	Name          string               `db:"name" json:"name"`
	Description   string               `db:"description" json:"description"`
	StorageMethod ProjectStorageMethod `db:"storage_method" json:"storage_method"`
	StorageSource []byte               `db:"storage_source" json:"storage_source"`
	ImportJobID   uuid.UUID            `db:"import_job_id" json:"import_job_id"`
}

func (q *sqlQuerier) InsertProjectHistory(ctx context.Context, arg InsertProjectHistoryParams) (ProjectHistory, error) {
	row := q.db.QueryRowContext(ctx, insertProjectHistory,
		arg.ID,
		arg.ProjectID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Description,
		arg.StorageMethod,
		arg.StorageSource,
		arg.ImportJobID,
	)
	var i ProjectHistory
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StorageMethod,
		&i.StorageSource,
		&i.ImportJobID,
	)
	return i, err
}

const insertProjectParameter = `-- name: InsertProjectParameter :one
INSERT INTO
  project_parameter (
    id,
    created_at,
    project_history_id,
    name,
    description,
    default_source,
    allow_override_source,
    default_destination,
    allow_override_destination,
    default_refresh,
    redisplay_value,
    validation_error,
    validation_condition,
    validation_type_system,
    validation_value_type
  )
VALUES
  (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14,
    $15
  ) RETURNING id, created_at, project_history_id, name, description, default_source, allow_override_source, default_destination, allow_override_destination, default_refresh, redisplay_value, validation_error, validation_condition, validation_type_system, validation_value_type
`

type InsertProjectParameterParams struct {
	ID                       uuid.UUID           `db:"id" json:"id"`
	CreatedAt                time.Time           `db:"created_at" json:"created_at"`
	ProjectHistoryID         uuid.UUID           `db:"project_history_id" json:"project_history_id"`
	Name                     string              `db:"name" json:"name"`
	Description              string              `db:"description" json:"description"`
	DefaultSource            sql.NullString      `db:"default_source" json:"default_source"`
	AllowOverrideSource      bool                `db:"allow_override_source" json:"allow_override_source"`
	DefaultDestination       sql.NullString      `db:"default_destination" json:"default_destination"`
	AllowOverrideDestination bool                `db:"allow_override_destination" json:"allow_override_destination"`
	DefaultRefresh           string              `db:"default_refresh" json:"default_refresh"`
	RedisplayValue           bool                `db:"redisplay_value" json:"redisplay_value"`
	ValidationError          string              `db:"validation_error" json:"validation_error"`
	ValidationCondition      string              `db:"validation_condition" json:"validation_condition"`
	ValidationTypeSystem     ParameterTypeSystem `db:"validation_type_system" json:"validation_type_system"`
	ValidationValueType      string              `db:"validation_value_type" json:"validation_value_type"`
}

func (q *sqlQuerier) InsertProjectParameter(ctx context.Context, arg InsertProjectParameterParams) (ProjectParameter, error) {
	row := q.db.QueryRowContext(ctx, insertProjectParameter,
		arg.ID,
		arg.CreatedAt,
		arg.ProjectHistoryID,
		arg.Name,
		arg.Description,
		arg.DefaultSource,
		arg.AllowOverrideSource,
		arg.DefaultDestination,
		arg.AllowOverrideDestination,
		arg.DefaultRefresh,
		arg.RedisplayValue,
		arg.ValidationError,
		arg.ValidationCondition,
		arg.ValidationTypeSystem,
		arg.ValidationValueType,
	)
	var i ProjectParameter
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.ProjectHistoryID,
		&i.Name,
		&i.Description,
		&i.DefaultSource,
		&i.AllowOverrideSource,
		&i.DefaultDestination,
		&i.AllowOverrideDestination,
		&i.DefaultRefresh,
		&i.RedisplayValue,
		&i.ValidationError,
		&i.ValidationCondition,
		&i.ValidationTypeSystem,
		&i.ValidationValueType,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO
  users (
    id,
    email,
    name,
    login_type,
    revoked,
    hashed_password,
    created_at,
    updated_at,
    username
  )
VALUES
  ($1, $2, $3, $4, false, $5, $6, $7, $8) RETURNING id, email, name, revoked, login_type, hashed_password, created_at, updated_at, temporary_password, avatar_hash, ssh_key_regenerated_at, username, dotfiles_git_uri, roles, status, relatime, gpg_key_regenerated_at, _decomissioned, shell
`

type InsertUserParams struct {
	ID             string    `db:"id" json:"id"`
	Email          string    `db:"email" json:"email"`
	Name           string    `db:"name" json:"name"`
	LoginType      LoginType `db:"login_type" json:"login_type"`
	HashedPassword []byte    `db:"hashed_password" json:"hashed_password"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	Username       string    `db:"username" json:"username"`
}

func (q *sqlQuerier) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser,
		arg.ID,
		arg.Email,
		arg.Name,
		arg.LoginType,
		arg.HashedPassword,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Name,
		&i.Revoked,
		&i.LoginType,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TemporaryPassword,
		&i.AvatarHash,
		&i.SshKeyRegeneratedAt,
		&i.Username,
		&i.DotfilesGitUri,
		pq.Array(&i.Roles),
		&i.Status,
		&i.Relatime,
		&i.GpgKeyRegeneratedAt,
		&i.Decomissioned,
		&i.Shell,
	)
	return i, err
}

const updateAPIKeyByID = `-- name: UpdateAPIKeyByID :exec
UPDATE
  api_keys
SET
  last_used = $2,
  expires_at = $3,
  oidc_access_token = $4,
  oidc_refresh_token = $5,
  oidc_expiry = $6
WHERE
  id = $1
`

type UpdateAPIKeyByIDParams struct {
	ID               string    `db:"id" json:"id"`
	LastUsed         time.Time `db:"last_used" json:"last_used"`
	ExpiresAt        time.Time `db:"expires_at" json:"expires_at"`
	OIDCAccessToken  string    `db:"oidc_access_token" json:"oidc_access_token"`
	OIDCRefreshToken string    `db:"oidc_refresh_token" json:"oidc_refresh_token"`
	OIDCExpiry       time.Time `db:"oidc_expiry" json:"oidc_expiry"`
}

func (q *sqlQuerier) UpdateAPIKeyByID(ctx context.Context, arg UpdateAPIKeyByIDParams) error {
	_, err := q.db.ExecContext(ctx, updateAPIKeyByID,
		arg.ID,
		arg.LastUsed,
		arg.ExpiresAt,
		arg.OIDCAccessToken,
		arg.OIDCRefreshToken,
		arg.OIDCExpiry,
	)
	return err
}
