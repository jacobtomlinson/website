---
author: Jacob Tomlinson
title: Why taking notes still matters
date: 2026-05-11T00:00:00+00:00
draft: false
categories:
  - blog
tags:
  - Notes
  - Work
  - Obsidian
  - AI
---

AI tools have radically changed note taking. All you need to do is record/transcribe your meeting and your favourite AI tool will write your notes for you. Magic! As a result I've noticed a lot of people stop taking notes altogether and delegating the whole thing to tools like Copilot. After all, why write notes if the AI generated notes in Microsoft Teams are going to be more comrehensive and complete with no effort required on your part?

As someone who discovered the values of note taking later in life I wanted to write a post on how I take notes and why I still take notes by hand. When I was in school I didn't bother with notes much becuase I always thought "it's all in the textbook anyway", I didn't understand the value of writing things down myself, and writing things in a way that suited my brain. With people moving to AI note takers it feels like the "it's all in the textbook" approach is being adopted en-masse, but where the textbook for every topic writes itself automatically. Textbooks are great, but notes are great too!

## How I take notes

I've been taking notes in [Obsidian](https://obsidian.md/) for a little over a year now. My workflow is very simple and I like it that way. There are thousands of posts out there telling you all the best workflows, all the best plugins and how to build a bespoke supercharged second brain. You can spend hours and hours tweaking your workflow and installing tools only to abandon it a few weeks or months later. You don't need all that complexity.

Looking through my settings I have a few plugins installed, but most are just small quality of life tweaks. There is only one community plugin that is core to my workflow and that's [Rollover Daily Todos](https://github.com/lumoe/obsidian-rollover-daily-todos).

I love the [daily note feature in Obsidian](https://obsidian.md/help/plugins/daily-notes). I treat this note as my home page. Every time I open Obsidian I click the "Open today's daily note" button, it's become a habit, it's my main starting point. My daily note template is also pretty simple, it looks something like this:

```markdown
# {{date:dddd, MMMM, Do YYYY}}

---

## Daily Log
- 

## TODO
### Morning routine
- [ ] Organise TODO lists
- [ ] Housekeep yesterday's daily note
- [ ] Inbox Zero
- [ ] Check GitHub Notifications

### Today
- [ ] 

### This week
- [ ] 

## Meetings
- 

## Notes
- 
```

The _daily log_ section is just a list of things I've done that day. I update this periodically throughout the day. It's usually things like "Went for a swim before work", "Made a Pull Request to fix issue org/project#id" and "Met with the team to discuss widgets". This is my very loose form of time tracking and accountability, it makes it easy for me to share what I've done during standups, in monthly reports or in my annual performance evaluation. These notes are easy to scroll through and are the breadcrumbs that on aggregate describe what I've achieved.

Next is my _TODO_ section. I break this down into daily repetitive tasks that are included automatically in the [template](https://obsidian.md/help/plugins/templates), tasks I want to do today and tasks I want to do this week. This is where [Rollover Daily Todos](https://github.com/lumoe/obsidian-rollover-daily-todos) comes in, whenever my next daily note is created it moves all the unchecked items from the previous daily note and puts them in the "This week" section. I do this because I find digital TODO lists problematic, if all of my TODO items aren't clearly visible in one place then they become out-of-sight out-of-mind. I used to use the Apple Reminders app for task tracking, and Google Keep before that. I had hierarchical sets of TODOs all layed out nicely but I would just forget about them. I got married in 2017 and I still have a "Wedding" category in Apple Reminders with unchecked items on it that I stumble across occasionally.

With my Obsidian workflow I know that all of my TODO items are available in a single note, today's note, and I can open that list by clicking "Open today's daily note". That's why it feels like a home page, everything hangs off this chronological capturing of TODOs. If I look through old daily notes I can see all the items I checked off that day, and that's it. As everything gets dumped into the "This week" section by the plugin I spend a little time each day choosing items I want to move up to the "Today" section. I also review the items and ask myself if I will really get it done this week. If I wont I probably need to move it off my TODO list to some other bigger picture tracking like a team GitHub project board, or delete it all together. 

As I check off items throughout the day I often add a corresponding item to my daily log. I think it's fine to duplicate these as they serve difference purposes, one looks forward, the other is a retrospective. Sometimes many TODO items will be condensed into a single daily log item.

Next I have a _Meeting Notes_ section. This is a simple bullet list. For each meeting I have I add an item, and then indent it to add notes on that specific meeting. I start every set of meeting notes this way because it is very low effort, I don't need to think about where meeting notes will go. I always have Obsi.dian open on my second screen, and I'm usually on my daily note already. So I just throw in a quick meeting title and start taking bulleted notes. If the notes stay at a handful of items then I will just leave this as it is, especially if it's just a 1:1 chat with someone. However, if it's a bigger meeting and the notes start growing and getting multiple levels of indentation I will extract them into a separate note. Sometimes the meeting section in my daily note just becomes a small bullet list of links to full meeting notes, other times some notes get left in the daily note, it doesn't matter. It's more important to me that something is captured somewhere in a low friction way. As soon as the barrier to taking notes becomes too high I wont do it, so low-friction is key.

Last I have my general _Notes_ section. This is also just a bullet list and behaves similarly to the meeting notes. It's the catchall place that is accessible in a click or two where I can jot things down. I'm not precious about this. If someone sends me a link to a cool GitHub repo on Slack I'll add a bullet with the link and a quick note on what it is. If I'm debugging something and I want to capture steps or other info I'll throw a bullet in and nest the notes underneath. If any of these grow beyond a few lines I'll extract them to a new note and link it from my daily note.

Almost every note in my Obsidian vault is referenced by at least one daily note, the one that triggered it's creation. 

Doing things this way means my Obsidian Vault's folder structure is pretty simple:

```text
├── Bases
├── Daily
│   ├── 2025
│   │   ├── 02-February
│   │   ├── ...
│   │   └── 12-December
│   └── 2026
│       ├── 01-January
│       ├── ...
│       └── 05-May
├── Files
├── Notes
└── Templates
```

The daily notes plugin puts my dailies into a `Daily/Year/Month` structure to make it easier for me to find a specific note. Although I do quite like the [Calendar](https://github.com/liamcain/obsidian-calendar-plugin) plugin for this and rarely use the folder structure to navigate them these days.

All of my other notes just go into a big `Notes` folder. I never navigate them this way, I always use [cross referencing](https://obsidian.md/help/links) or search to find notes, so they may as well just go into one blob. People waste way too much time trying to come up with hierarchies and ontologies to categorise and sort their notes. I much prefer tags and metadata for this.

I also have a `Files` folder which is just the default place where images and photos go. Usually the only things that get put in here are screenshots that I paste into notes, or photos of slides if I'm taking notes on a talk. 

I have a couple of [Obsidian Bases](https://obsidian.md/help/bases), mostly for things like people or books. It can help to see all this metadata in one place and bases make that easier, but it's by no means core to my workflow.

Last is a `Templates` directory. I have a bunch of templates for things like my daily note, meeting notes, talks, trips, etc. These are just loose hierarchies of structure that I commonly use. Again the only template that's actually important to me is the daily note template, everything else is just polish and time saving.

## The value in typing your notes

All of these notes, everything that I write in my daily note, and everything I write in these extracted notes is written by me. I feel strongly about this because I find that at least half of the value in my notes is the act of taking them. 

If I attend a talk, watch an educational YouTube video or attend a meeting I take notes. I write down the things I find interesting and the things that are important to me. I paraphrase and restate the knowledge from the session in my own words. Doing this causes me to actively listen and pay attention. It can be too tempting to multitask and keep writing code during a meeting, especially when you work from home. By taking notes it forces me to focus on the task at hand, and I get a lot of value from this discipline. I hear more and retain more. 

The act of writing out the same information that I just heard also causes me to process the information more deeply. My mind doesn't wander, I dont just tune out things I don't understand. I can write down questions I have and track things I want to look up later. Honestly I would still take notes even if they got deleted after each meeting, just the act of writing things down gives me a lot of value.

After a meeting I am left with a note that contains my perspective on the meeting. It is a filtered and incomplete view of what was discussed. It contains my interpretation of what was said, points that were interesting to me and any embellishments I've made to the note. I find I often enrich notes by adding context based on things I know from outside the meeting, I may cross link topics to other notes and look things up as the person is talking and include other notes, links or context. Anything I think I might find useful later.

If everyone in the meeting all took notes they would be very different. Each would be an incomplete view of the meeting from that person's perspective. They would probably contain more information than was actually discussed in the meeting. More on that later.

## Notes as a resource

Over time these notes build up a corpus of knowledge. This is the second brain that people spend so much time talking about on YouTube. This is also something I find useful. I often search through my notes to remind myself of things that were discussed, or to view relationships between things. For instance I may hear about a new project in a meeting, then weeks later that project comes up again, it's easy for me to find which meeting it was discussed and then figure out who to ask about it based on the attendee list.

This is where AI note takers can be powerful. They capture every meeting, all of the things that were said. They don't capture my narrow view, they capture an objectively complete view. Being able to search through meetings with AI is an excellent way to recall things. When you combine that with other knowledge sources like Slack, GitHub and your documents you can have an AI chat that helps you retrieve almost any information you have come into contact with. I love being able to ask an AI Agent questions like "Who did I talk to recently that mentioned <insert topic here>? They mentioned a URL to a documentation page, can you find it?". 

While AI notes are complete transcriptons and summaries of meetings my hand written notes are opinionated views onto that same data, with topics that are important to me captured more prominently. I tend to think of my files as "Notes" and AI files as "Logs". Both are valuable and complementary. Going forwards my plan is to automate more of my note taking to aid this retrieval, but without replacing what I am currently doing. For each meeting note that I write by hand I want to cross reference the AI generated version, so that I can search both simultaneously.

## Taking notes as a group

Of course not all of my notes live in my Obsidian Vault. A handful of meetings I attend have shared notes, like our weekly team meeting. These notes may live in a Google Doc that is linked from the meeting invite and editible by anyone in the team. In my Obsidian notes I usually have a corresponding note that just contains a link to the Google Doc to allow me to cross-reference it from other notes. This does make retrieval a little trickier, I wish all the text was on my local disk for easy searching, but I also care about being flexible and it's more valuable to the team keep these notes in a central place.

One thing I love about taking notes as a team is that other social dynamics come into play. This is added value that can't be given by AI. By all means use AI to transcribe your meetings, and link those autogenerated notes from your team notes, but you should still take notes by hand as a team too.

A team standup may involve everyone taking it in turns to go around the room and share what you've done, or what you are working on now. It's almost impossible to write down what you are doing while also speaking out loud, so having someone else writing down what you are saying is very helpful. But where this gets interesting is when a human is paraphrasing what you are saying because while you can't really paraphrase yourself while speaking simultaneously we are very good at scanning the notes as they are being written. This allows you to see their comprehension of what you said and course correct automatically as you are speaking. If someone is writing down what I am saying in their own words, I can get a good gauge on whether they are taking the meaning I am intending. If they write down something that doesn't quite capture what I intended to say I can easily follow up or repeat myself to make sure I am well understood.

I care about other people's comprehension of what I am saying in meetings, it's the whole point of the meeting. By having those people take notes while I am talking I can fine-tune their comprehension on the fly. The AI notes then complement this because you also have a literal representation of what I said. The human note takers may have added links or information based on their own knowledge, they likely enriched the notes as they wrote them. So when it comes to retriavel later via an LLM the combination of enriched, course-corrected human notes and literal transcription-based notes means you can retieve information that is greater than the sum of it's parts.

## Conclusion

I love AI tools. The explosion of powerful models in the last five years has made me more excited about technology than ever, but I am apprehensive about how much we are delegating to them and what we are losing in the process. 

A decent speach-to-text model combined with an LLM can take way more comprehensive notes than I can, it can capture everything that was said, it can paraphrase, it can create high-level summaries and it can do all of this without any effort on my part. This gives me huge value.

Actively listening, writing down my own comprehension of a topic and enriching what was said with other knowledge that I have is also extremely important. I use manual note taking as a mechanism to make myself pay attention and capture opinionated knowledge. This gives me different value.

Having the notes and the textbook combined is the best of both worlds. I wish someone took the time to explain that to me when I was 16.
