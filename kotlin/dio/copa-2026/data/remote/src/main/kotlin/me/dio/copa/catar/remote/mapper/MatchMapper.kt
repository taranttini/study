package me.dio.copa.catar.remote.mapper

import me.dio.copa.catar.domain.model.MatchDomain
import me.dio.copa.catar.domain.model.StadiumDomain
import me.dio.copa.catar.domain.model.Team
import me.dio.copa.catar.domain.model.TeamFlag
import me.dio.copa.catar.remote.model.MatchRemote
import me.dio.copa.catar.remote.model.StadiumRemote
import java.time.LocalDateTime
import java.time.ZoneId
import java.util.Date
import java.util.Locale

internal fun List<MatchRemote>.toDomain() = map { it.toDomain() }

fun MatchRemote.toDomain(): MatchDomain {
    return MatchDomain(
        id = "$team1-$team2",
        name = name,
        team1 = team1.toTeam(),
        team1_flag = team1_flag.toFlag(),
        team2 = team2.toTeam(),
        team2_flag = team2_flag.toFlag(),
        stadium = stadium.toDomain(),
        date = date.toLocalDateTime(),
    )
}

private fun Date.toLocalDateTime(): LocalDateTime {
    return toInstant().atZone(ZoneId.systemDefault()).toLocalDateTime()
}

private fun String.toTeam(): Team {
    return Team(
        displayName = this.toString()
    )
}

private fun String.toFlag(): TeamFlag {
    return TeamFlag(
        flagUrl = getTeamFlag(this),
    )
}

private fun getTeamFlag(team: String): String {
    return team
}

fun StadiumRemote.toDomain(): StadiumDomain {
    return StadiumDomain(
        name = name,
        image = image
    )
}
