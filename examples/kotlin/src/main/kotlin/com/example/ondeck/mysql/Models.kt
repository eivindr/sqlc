// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package com.example.ondeck.mysql

import java.sql.Timestamp
import java.time.Instant

enum class VenuesStatus(val value: String) {
  OPEN("open"),
  CLOSED("closed");

  companion object {
    private val map = VenuesStatus.values().associateBy(VenuesStatus::value)
    fun lookup(value: String) = map[value]
  }
}

data class City (
  val slug: String,
  val name: String
)

// Venues are places where muisc happens
data class Venue (
  val id: Long,
  // Venues can be either open or closed
  val status: VenuesStatus,
  val statuses: String?,
  // This value appears in public URLs
  val slug: String,
  val name: String,
  val city: String,
  val spotifyPlaylist: String,
  val songkickId: String?,
  val tags: String?,
  val createdAt: Instant
)

