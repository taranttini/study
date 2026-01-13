package me.dio.copa.catar.features

import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material.Card
import androidx.compose.material.MaterialTheme
import androidx.compose.material.Surface
import androidx.compose.material.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.alpha
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import me.dio.copa.catar.R
import me.dio.copa.catar.domain.extensions.getDate
import me.dio.copa.catar.domain.model.MatchDomain
import me.dio.copa.catar.domain.model.TeamDomain
import me.dio.copa.catar.domain.model.TeamFlag
import me.dio.copa.catar.ui.theme.Shapes
import okhttp3.internal.wait

typealias NotificationOnClick = (match: MatchDomain) -> Unit

@Composable
fun MainScreen(matches: List<MatchDomain>, onNotificationClick: NotificationOnClick) {
    Box(
        modifier = Modifier
            .fillMaxSize()
            .padding(16.dp)
    ) {
        LazyColumn(verticalArrangement = Arrangement.spacedBy(8.dp)) {
            items(matches) { match ->
                MatchInfo(match, onNotificationClick)
            }
        }
    }
}

@Composable
fun MatchInfo(match: MatchDomain, onNotificationClick: NotificationOnClick) {
    Card(
        shape = Shapes.large,
        modifier = Modifier.fillMaxWidth()
    ) {
        Box {
            AsyncImage(
                model = match.stadium.image,
                contentDescription = null,
                contentScale = ContentScale.Crop,
                modifier = Modifier.height(160.dp)
            )

            Column(modifier = Modifier.padding(16.dp)) {
                Notification(match, onNotificationClick)
                Title(match)
                Teams(match)
            }
        }
    }
}

@Composable
fun Notification(match: MatchDomain, onClick: NotificationOnClick) {
    Row(modifier = Modifier.fillMaxWidth(), horizontalArrangement = Arrangement.End) {
        val drawable = if (match.notificationEnabled) R.drawable.ic_notifications_active
        else R.drawable.ic_notifications

        Image(
            painter = painterResource(id = drawable),
            modifier = Modifier.clickable {
                onClick(match)
            },
            contentDescription = null
        )
    }
}

@Composable
fun Title(match: MatchDomain) {
    Row(
        modifier = Modifier.fillMaxWidth(),
        horizontalArrangement = Arrangement.Center
    ) {
        Surface(
            color = Color.Black.copy(alpha = .5f),
            modifier = Modifier.fillMaxSize()
        ) {
            Text(
                text = "${match.date.getDate()} - ${match.name}",
                textAlign = TextAlign.Center,
                style = MaterialTheme.typography.h6.copy(color = Color.White),
                modifier = Modifier
                    .background(Color.Transparent)
                    .fillMaxSize()

            )
        }

    }
}

@Composable
fun Teams(match: MatchDomain) {

    Column() {
        Spacer(modifier = Modifier.size(10.dp))

        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.Center,
            verticalAlignment = Alignment.CenterVertically,
        ) {
            TeamItem(
                team = match.team1,
                teamFlag = match.team1_flag,
                modifier = Modifier.weight(0.5f)
            )

            Text(
                text = "X",
                modifier = Modifier.padding(end = 16.dp, start = 16.dp),
                style = MaterialTheme.typography.h6.copy(color = Color.Gray)
            )

            TeamItem(
                team = match.team2,
                teamFlag = match.team2_flag,
                modifier = Modifier.weight(0.5f)
            )
        }

    }
}

@Composable
fun TeamItem(team: TeamDomain, teamFlag: TeamFlag, modifier: Modifier = Modifier) {
    Column(
        horizontalAlignment = Alignment.CenterHorizontally,
        modifier = modifier
    ) {
        Spacer(modifier = Modifier.size(10.dp))

        AsyncImage(
            model = teamFlag.flagUrl,
            contentDescription = "",
            contentScale = ContentScale.Fit,
            modifier = Modifier.height(26.dp)
        )

        Spacer(modifier = Modifier.size(4.dp))

        Surface(
            color = Color.Black.copy(alpha = .3f),
            modifier = Modifier.fillMaxSize()
        ) {
            Text(
                text = team.displayName,
                textAlign = TextAlign.Center,
                style = MaterialTheme.typography.body1.copy(color = Color.White)
            )
        }
    }
}
